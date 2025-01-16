import 'package:flutter/material.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/models/user_app.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/providers/connectivity_provider.dart';
import 'package:squad_go/core/services/team_service.dart';
import 'package:squad_go/main.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/add_player.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/add_team.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/edit_player.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/edit_stats_player.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/edit_team.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/offline.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/show_player_details.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class TeamsHandle extends StatefulWidget {
  final String eventId;
  final String? eventCreatorId;
  final List<Team> teams;
  final int maxTeams;
  final bool canEdit;
  final Color color;
  final bool isEventFinished;
  final bool isEventNowPlaying;
  final Future<void> Function()? onRefresh;

  const TeamsHandle({
    super.key,
    required this.teams,
    required this.eventId,
    required this.canEdit,
    required this.maxTeams,
    required this.color,
    this.isEventFinished = false,
    this.isEventNowPlaying = false,
    this.onRefresh,
    this.eventCreatorId,
  });

  @override
  State<TeamsHandle> createState() => _TeamsHandleState();
}

class _TeamsHandleState extends State<TeamsHandle> {
  final TeamService teamService = TeamService();

  UserApp? currentUser;
  bool userHasTeam = false;

  @override
  void initState() {
    super.initState();

    currentUser = context.read<AuthState>().userInfo;
  }

  void _joinTeam(String teamId, String teamName) async {
    final translate = AppLocalizations.of(context);
    _showAlertDialog(
      '${translate?.join_button ?? "Rejoindre"} $teamName',
      translate?.confirm_join_team ??
          "Voulez-vous vraiment rejoindre cette équipe?",
      () async {
        try {
          await teamService.joinTeam(widget.eventId, teamId);
          widget.onRefresh?.call();
        } catch (e) {
          // Handle error
          log.severe('Failed to join team: $e');
        }
      },
    );
  }

  void _switchTeam(String teamId, String teamName) async {
    final translate = AppLocalizations.of(context);
    _showAlertDialog(
      translate?.change_team ?? "Changer d'équipe",
      '${translate?.confirm_change_team ?? "Voulez-vous vraiment changer d\'équipe et rejoindre"} $teamName ?',
      () async {
        try {
          await teamService.switchTeam(widget.eventId, teamId);
          widget.onRefresh?.call();
        } catch (e) {
          // Handle error
          log.severe('Failed to switch team: $e');
        }
      },
    );
  }

  void _showAlertDialog(String title, String content, void Function() onOk) {
    final translate = AppLocalizations.of(context);
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: Text(title),
          content: Text(content),
          actions: [
            TextButton(
              child: Text(translate?.cancel ?? "Annuler"),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
            TextButton(
              child: Text("OK"),
              onPressed: () {
                onOk();
                Navigator.of(context).pop();
              },
            ),
          ],
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);
    userHasTeam = widget.teams.any(
      (team) => team.players.any((player) => player.userID == currentUser!.id),
    );

    var isOnline = context.watch<ConnectivityState>().isConnected;

    return DefaultTabController(
      length: widget.teams.length,
      child: Column(
        children: [
          Stack(
            children: [
              SizedBox(
                height: 40,
                child: TabBar(
                  isScrollable: true,
                  padding: EdgeInsets.only(right: 100),
                  dividerColor: Colors.transparent,
                  indicatorSize: TabBarIndicatorSize.tab,
                  tabAlignment: TabAlignment.start,
                  indicator: BoxDecoration(
                    color: Colors.blue,
                    borderRadius: BorderRadius.circular(10),
                  ),
                  labelColor: Colors.white,
                  labelStyle: TextStyle(fontWeight: FontWeight.bold),
                  labelPadding: EdgeInsets.only(left: 16),
                  tabs: widget.teams.map(
                    (team) {
                      return Tab(
                        child: Row(
                          children: [
                            Text(team.name),
                            SizedBox(width: widget.canEdit ? 8 : 16),
                            if (widget.canEdit && !widget.isEventFinished) ...[
                              Container(
                                width: 30,
                                height: 30,
                                decoration: BoxDecoration(
                                  shape: BoxShape.circle,
                                  color: Colors.blueAccent,
                                ),
                                child: IconButton(
                                  icon: Icon(
                                    Icons.edit,
                                    size: 16,
                                    color: Colors.white,
                                  ),
                                  padding: EdgeInsets.zero,
                                  onPressed: () {
                                    if (!isOnline) {
                                      showDialog(
                                        context: context,
                                        builder: (context) =>
                                            const OfflineDialog(),
                                      );
                                      return;
                                    }
                                    showDialog(
                                      context: context,
                                      builder: (context) {
                                        return EditTeamDialog(
                                          eventId: widget.eventId,
                                          team: team,
                                          onRefresh: widget.onRefresh,
                                        );
                                      },
                                    );
                                  },
                                ),
                              ),
                              SizedBox(width: 8),
                            ],
                          ],
                        ),
                      );
                    },
                  ).toList(),
                ),
              ),
              if (widget.canEdit &&
                  !widget.isEventFinished &&
                  (widget.maxTeams == 0 ||
                      widget.teams.length < widget.maxTeams))
                Positioned(
                  right: 0,
                  child: SizedBox(
                    height: 40,
                    child: ElevatedButton(
                      onPressed: () {
                        if (!isOnline) {
                          showDialog(
                            context: context,
                            builder: (context) => const OfflineDialog(),
                          );
                          return;
                        }
                        showDialog(
                          context: context,
                          builder: (context) {
                            return AddTeamDialog(
                              eventId: widget.eventId,
                              onRefresh: widget.onRefresh,
                            );
                          },
                        );
                      },
                      style: ElevatedButton.styleFrom(
                        shape: CircleBorder(),
                        foregroundColor: Colors.white,
                        backgroundColor: Colors.blue,
                      ),
                      child: Icon(
                        Icons.add,
                        color: Colors.white,
                      ),
                    ),
                  ),
                )
            ],
          ),
          Expanded(
            child: TabBarView(
              children: widget.teams
                  .map(
                    (team) => Column(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Row(
                          mainAxisAlignment: MainAxisAlignment.spaceBetween,
                          children: [
                            if (!widget.isEventFinished &&
                                (!userHasTeam ||
                                    (userHasTeam &&
                                        !team.players.any((player) =>
                                            player.userID == currentUser!.id))))
                              TextButton(
                                onPressed: () {
                                  if (!isOnline) {
                                    showDialog(
                                      context: context,
                                      builder: (context) =>
                                          const OfflineDialog(),
                                    );
                                    return;
                                  }
                                  if (userHasTeam) {
                                    _switchTeam(team.id, team.name);
                                  } else {
                                    _joinTeam(team.id, team.name);
                                  }
                                },
                                child: Badge(
                                  label: Row(
                                    children: [
                                      Icon(
                                        userHasTeam
                                            ? Icons.swap_horiz
                                            : Icons.group_add,
                                        size: 16,
                                        color: Colors.white,
                                      ),
                                      SizedBox(width: 6),
                                      Text(
                                        userHasTeam
                                            ? translate?.change_team ??
                                                "Changer d'équipe"
                                            : translate?.join_team ??
                                                "Rejoindre l'équipe",
                                      ),
                                    ],
                                  ),
                                  backgroundColor:
                                      Theme.of(context).colorScheme.primary,
                                  textStyle: TextStyle(
                                    fontWeight: FontWeight.bold,
                                  ),
                                  padding: EdgeInsets.symmetric(
                                    horizontal: 7,
                                    vertical: 5,
                                  ),
                                ),
                              ),
                            if (widget.isEventFinished ||
                                (userHasTeam &&
                                    team.players.any((player) =>
                                        player.userID == currentUser!.id)))
                              TextButton(
                                onPressed: null,
                                child: SizedBox.shrink(),
                              ),
                            Badge(
                              label: Text(
                                  '${team.players.length}${team.maxPlayers > 0 ? ' / ${team.maxPlayers}' : ''}'),
                              backgroundColor: team.maxPlayers == 0
                                  ? Colors.grey.shade400
                                  : team.players.length < team.maxPlayers
                                      ? Colors.green.shade300
                                      : Colors.red.shade300,
                              textStyle: TextStyle(
                                fontWeight: FontWeight.bold,
                              ),
                              padding: EdgeInsets.symmetric(
                                horizontal: 7,
                                vertical: 5,
                              ),
                            )
                          ],
                        ),
                        team.players.isNotEmpty
                            ? Expanded(
                                child: ListView.builder(
                                  itemCount: team.players.length,
                                  itemBuilder: (context, index) {
                                    final player = team.players[index];
                                    final isCurrentUser =
                                        player.userID == currentUser!.id;
                                    final isEventCreator =
                                        widget.eventCreatorId == player.userID;
                                    final Color roleColor =
                                        player.role == PlayerRole.coach
                                            ? Colors.blue
                                            : Colors.deepOrange;
                                    return Container(
                                      margin: EdgeInsets.only(bottom: 16),
                                      decoration: BoxDecoration(
                                        color: isCurrentUser
                                            ? widget.color.withAlpha(40)
                                            : widget.color.withAlpha(20),
                                        borderRadius: BorderRadius.circular(16),
                                      ),
                                      child: InkWell(
                                        onTap: () {
                                          showDialog(
                                            context: context,
                                            builder: (context) {
                                              return ShowPlayerDetailsDialog(
                                                eventId: widget.eventId,
                                                player: player,
                                                canEdit: widget.canEdit,
                                                isCurrentUser: isCurrentUser,
                                                onRefresh: widget.onRefresh,
                                              );
                                            },
                                          );
                                        },
                                        child: ListTile(
                                          title: Row(
                                            children: [
                                              Text(
                                                player.name ?? player.email,
                                                style: TextStyle(
                                                  fontWeight: isCurrentUser
                                                      ? FontWeight.bold
                                                      : FontWeight.normal,
                                                  color: isCurrentUser
                                                      ? widget.color
                                                      : null,
                                                ),
                                              ),
                                              if (isEventCreator)
                                                Padding(
                                                  padding:
                                                      const EdgeInsets.only(
                                                          left: 8.0),
                                                  child: Icon(
                                                    Icons.star,
                                                    color: Colors.amber,
                                                    size: 16,
                                                  ),
                                                ),
                                            ],
                                          ),
                                          subtitle: player.role !=
                                                  PlayerRole.player
                                              ? Row(
                                                  mainAxisAlignment:
                                                      MainAxisAlignment.start,
                                                  children: [
                                                    Badge(
                                                      label: Text(
                                                        player.role ==
                                                                PlayerRole.coach
                                                            ? 'Coach'
                                                            : translate
                                                                    ?.organizer ??
                                                                'Organisateur',
                                                        style: TextStyle(
                                                          color: roleColor,
                                                          fontWeight:
                                                              FontWeight.bold,
                                                        ),
                                                      ),
                                                      backgroundColor: roleColor
                                                          .withAlpha(40),
                                                      padding:
                                                          EdgeInsets.symmetric(
                                                        horizontal: 5,
                                                        vertical: 2,
                                                      ),
                                                    ),
                                                  ],
                                                )
                                              : null,
                                          trailing: player.status ==
                                                  PlayerStatus.valid
                                              ? Row(
                                                  mainAxisSize:
                                                      MainAxisSize.min,
                                                  children: [
                                                    if (widget.canEdit &&
                                                        !widget
                                                            .isEventFinished) ...[
                                                      IconButton(
                                                        icon: Icon(Icons.edit),
                                                        color: isCurrentUser
                                                            ? widget.color
                                                            : null,
                                                        onPressed: () {
                                                          if (!isOnline) {
                                                            showDialog(
                                                              context: context,
                                                              builder: (context) =>
                                                                  const OfflineDialog(),
                                                            );
                                                            return;
                                                          }
                                                          showDialog(
                                                            context: context,
                                                            builder: (context) {
                                                              return EditPlayerDialog(
                                                                eventId: widget
                                                                    .eventId,
                                                                player: player,
                                                                teams: widget
                                                                    .teams,
                                                                onRefresh: widget
                                                                    .onRefresh,
                                                              );
                                                            },
                                                          );
                                                        },
                                                      ),
                                                      if (widget
                                                              .isEventNowPlaying &&
                                                          player.role !=
                                                              PlayerRole.coach)
                                                        IconButton(
                                                          icon: Icon(
                                                              Icons.bar_chart),
                                                          color: isCurrentUser
                                                              ? widget.color
                                                              : null,
                                                          onPressed: () {
                                                            showDialog(
                                                              context: context,
                                                              builder:
                                                                  (context) {
                                                                return EditStatsPlayerDialog(
                                                                  eventId: widget
                                                                      .eventId,
                                                                  player:
                                                                      player,
                                                                  onRefresh: widget
                                                                      .onRefresh,
                                                                );
                                                              },
                                                            );
                                                          },
                                                        ),
                                                    ],
                                                    if (!widget.canEdit ||
                                                        widget.isEventFinished)
                                                      IconButton(
                                                        icon: Icon(
                                                          Icons.remove_red_eye,
                                                        ),
                                                        color: isCurrentUser
                                                            ? widget.color
                                                            : null,
                                                        onPressed: () {
                                                          showDialog(
                                                            context: context,
                                                            builder: (context) {
                                                              return ShowPlayerDetailsDialog(
                                                                eventId: widget
                                                                    .eventId,
                                                                player: player,
                                                                canEdit: widget
                                                                    .canEdit,
                                                                isCurrentUser:
                                                                    isCurrentUser,
                                                                onRefresh: widget
                                                                    .onRefresh,
                                                              );
                                                            },
                                                          );
                                                        },
                                                      ),
                                                  ],
                                                )
                                              : Badge(
                                                  label: Text(
                                                    translate?.waiting ??
                                                        'En attente',
                                                    style: TextStyle(
                                                      color: Colors.white,
                                                      fontWeight:
                                                          FontWeight.bold,
                                                    ),
                                                  ),
                                                  backgroundColor: Colors.grey,
                                                  padding: EdgeInsets.symmetric(
                                                    horizontal: 7,
                                                    vertical: 5,
                                                  ),
                                                ),
                                        ),
                                      ),
                                    );
                                  },
                                ),
                              )
                            : Expanded(
                                child: Container(
                                  margin: EdgeInsets.only(top: 16),
                                  child: Text(translate?.no_players_team ??
                                      "Actuellement aucun joueur dans cette équipe"),
                                ),
                              ),
                        if (widget.canEdit &&
                            !widget.isEventFinished &&
                            (team.maxPlayers == 0 ||
                                team.players.length < team.maxPlayers)) ...[
                          SizedBox(height: 8),
                          ElevatedButton.icon(
                            onPressed: () {
                              if (!isOnline) {
                                showDialog(
                                  context: context,
                                  builder: (context) => const OfflineDialog(),
                                );
                                return;
                              }
                              showDialog(
                                context: context,
                                builder: (context) {
                                  return AddPlayerDialog(
                                    eventId: widget.eventId,
                                    teamId: team.id,
                                    onRefresh: widget.onRefresh,
                                  );
                                },
                              );
                            },
                            icon: const Icon(
                              Icons.add,
                              color: Colors.white,
                            ),
                            label: Text(
                                translate?.add_player ?? 'Ajouter un joueur'),
                            style: ElevatedButton.styleFrom(
                              foregroundColor: Colors.white,
                              backgroundColor: Colors.blue,
                              padding: const EdgeInsets.symmetric(
                                  horizontal: 20, vertical: 10),
                              textStyle: const TextStyle(
                                fontSize: 16,
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                          ),
                        ],
                      ],
                    ),
                  )
                  .toList(),
            ),
          ),
        ],
      ),
    );
  }
}
