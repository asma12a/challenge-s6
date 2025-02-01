import 'dart:async';
import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/models/user_stats.dart';
import 'package:squad_go/core/services/sport_stat_labels_service.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class Score extends StatefulWidget {
  final List<Team> teams;
  final String eventId;
  final String sportId;
  final bool isEventNowPlaying;  // Nouvelle prop pour savoir si l'événement est en cours
  final bool isEventFinished;

  const Score({super.key, required this.teams, required this.eventId, required this.sportId, required this.isEventNowPlaying, required this.isEventFinished});

  @override
  _ScoreState createState() => _ScoreState();
}

class _ScoreState extends State<Score> {
  late List<Team> teams;
  List<Map<String, dynamic>> performeurs = []; // userId + value (points)
  late Timer _timer;
  final SportStatLabelsService statLabelsService = SportStatLabelsService();
  late String mainLabel = "points";



  @override
  void initState() {
    super.initState();
    teams = widget.teams;
    _fetchMainLabel(widget.sportId);

    debugPrint("event commencé ${widget.isEventNowPlaying}");
    if (widget.isEventNowPlaying && !widget.isEventFinished ) {
      _timer = Timer.periodic(Duration(seconds: 2), _fetchUserStats);
    }else if (widget.isEventFinished){
      _fetchUserStatsOnce();
    }
  }

  @override
  void dispose() {
    if (widget.isEventNowPlaying) {
      _timer.cancel();
    }
    super.dispose();
  }

  void _fetchMainLabel(String sportId) async{
    try{
      final labels =
      await statLabelsService.getStatLabelsBySport(widget.sportId);

      if (labels.isNotEmpty) {
        setState(() {
          mainLabel = labels.first.label.split(" ").first.toLowerCase();
        });
      }
    }catch (error) {
      print("Erreur lors de la récupération des stats: $error");
    }
  }


  void _fetchUserStats(Timer timer) async {
    try {
      final List<UserStats> userStats = await statLabelsService.getAllTeamUserMainStatByEvent(widget.eventId); // Remplace "EVENT_ID" par l'ID réel de l'événement

      setState(() {
        performeurs = userStats.map((stat) => {
          'userId': stat.user?.id,
          'value': stat.value ?? 0,
        }).toList();
      });
    } catch (error) {
      print("Erreur lors de la récupération des stats: $error");
    }
  }

  void _fetchUserStatsOnce() async {
    try {
      final List<UserStats> userStats = await statLabelsService.getAllTeamUserMainStatByEvent(widget.eventId); // Remplace "EVENT_ID" par l'ID réel de l'événement

      setState(() {
        performeurs = userStats.map((stat) => {
          'userId': stat.user?.id,
          'value': stat.value ?? 0,
        }).toList();
      });
    } catch (error) {
      print("Erreur lors de la récupération des stats: $error");
    }
  }


  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);

    if (!widget.isEventNowPlaying && !widget.isEventFinished) {
      return Center(
        child: Text(
          translate?.event_not_started ?? "Revenez quand l'événement sera commencé.",
        ),
      );
    }
    Map<String, int> teamScores = {};
    for (var team in teams) {
      int teamPoints = 0;
      for (var player in team.players) {
        final perf = performeurs.firstWhere(
          (p) => p['userId'] == player.userID,
          orElse: () => {'value': 0},
        );
        teamPoints += (perf['value'] as num).toInt();
      }
      teamScores[team.id] = teamPoints;
    }

    // Trouver le score max
    int maxScore = teamScores.values.isNotEmpty ? teamScores.values.reduce((a, b) => a > b ? a : b) : 0;

    List<Team> winningTeams = teams.where((team) => (teamScores[team.id] ?? 0) == maxScore).toList();

    bool noWinner = winningTeams.length == teams.length && maxScore == 0;

    Team? winner = (winningTeams.length == 1) ? winningTeams.first : null;
    return Container(
      padding: const EdgeInsets.all(16.0),
      child: SingleChildScrollView(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: teams
              .map((team) => _buildTeamCard(context, team, winner != null && team.id == winner.id, teamScores[team.id] ?? 0))
              .toList(),
        ),
      ),
    );
  }

  Widget _buildTeamCard(BuildContext context, Team team, bool isWinner, int score) {
    final translate = AppLocalizations.of(context);
    return Card(
      color: isWinner ? Colors.green[400] : Colors.grey[200],
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text(
                  team.name,
                  style: TextStyle(
                    fontSize: 18,
                    fontWeight: FontWeight.bold,
                    color: isWinner ? Colors.white : Colors.black,
                  ),
                ),
                Row(
                  children: [
                    Text(
                      "${translate?.score ?? 'Score'}: $score",
                      style: TextStyle(
                        fontSize: 22,
                        fontWeight: FontWeight.bold,
                        color: isWinner ? Colors.white : Colors.black,
                      ),
                    ),
                    if (isWinner) ...[
                      const SizedBox(width: 8),
                      const Icon(Icons.emoji_events, color: Colors.amber, size: 28),
                    ],
                  ],
                ),
              ],
            ),
            const Divider(),
            const SizedBox(height: 8),
            Text(
              translate?.players ?? "Joueurs",
              style: TextStyle(
                fontSize: 16,
                fontWeight: FontWeight.bold,
                color: isWinner ? Colors.white70 : Colors.grey[700],
              ),
            ),
            const SizedBox(height: 8),
            ...team.players.map((player) => _buildPlayerRow(player, isWinner)).toList(),
          ],
        ),
      ),
    );
  }

  Widget _buildPlayerRow(Player player, bool isWinner) {
    final perf = performeurs.firstWhere(
      (p) => p['userId'] == player.userID,
      orElse: () => {'value': 0},
    );
    int points = perf['value'];

    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        Text(
          player.name ?? player.email,
          style: TextStyle(color: isWinner ? Colors.white : Colors.black),
        ),
        Text(
          "$points $mainLabel",
          style: TextStyle(color: isWinner ? Colors.white : Colors.black),
        ),
      ],
    );
  }
}
