import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/services/sport_stat_labels_service.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/models/user_stats.dart';
import 'package:squad_go/core/services/team_service.dart';

import '../../../../main.dart';

class ShowPlayerDetailsDialog extends StatefulWidget {
  final String eventId;
  final Player player;
  final bool canEdit;
  final bool isCurrentUser;
  final Future<void> Function()? onRefresh;

  const ShowPlayerDetailsDialog({
    super.key,
    required this.eventId,
    required this.player,
    required this.canEdit,
    required this.isCurrentUser,
    this.onRefresh,
  });

  @override
  State<ShowPlayerDetailsDialog> createState() =>
      _ShowPlayerDetailsDialogState();
}

class _ShowPlayerDetailsDialogState extends State<ShowPlayerDetailsDialog> {
  final SportStatLabelsService statLabelsService = SportStatLabelsService();
  final EventService eventService = EventService();
  Event? event;
  late Player _player;
  int nbEvents = 0;
  List<Map<String, dynamic>> stats = [];

  @override
  void initState() {
    super.initState();
    _player = widget.player;
    _fetchEventDetails();
  }

  void _fetchEventDetails() async {
    try {
      final eventDetails = await eventService.getEventById(widget.eventId);
      setState(() {
        event = eventDetails;
      });
      _loadUserPerformance();
    } catch (e) {
      // Handle error
      log.severe('Failed to fetch event details: $e');
    }
  }

  Future<void> _loadUserPerformance() async {
    try {
      if (event!.id == null) return;

      final userPerformances = await statLabelsService
          .getUserPerformanceBySport(event!.sport.id, _player.userID);
      setState(() {
        nbEvents = userPerformances.nbEvents;
        stats = userPerformances.stats
                ?.map((stat) => {
                      'stat_label': {'label': stat.stat?.label ?? ''},
                      'value': stat.value,
                    })
                .toList() ??
            [];
      });
    } catch (e) {
      log.severe('Failed to fetch user performance: $e');
    }
  }

  void _deletePlayer() async {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: Text('Supprimer le joueur'),
          content: Text(
              'Êtes-vous sûr de vouloir supprimer le joueur ${_player.name} ?'),
          actions: [
            TextButton(
              child: Text("Annuler"),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
            TextButton(
              child: Text("OK"),
              onPressed: () async {
                try {
                  await TeamService().deletePlayer(widget.eventId, _player.id);
                  widget.onRefresh?.call();
                } on AppException catch (e) {
                  // Handle AppException error
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(content: Text('Erreur: ${e.message}')),
                  );
                } catch (e) {
                  // Handle other errors
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(content: Text('Une erreur est survenue')),
                  );
                }
                Navigator.of(context).pop();
                Navigator.of(context).pop();
              },
            ),
          ],
        );
      },
    );
  }

  void _leaveTeam() async {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: Text('Quitter l\'équipe'),
          content: Text(
              'Êtes-vous sûr de vouloir quitter l\'équipe et l\'événement ?'),
          actions: [
            TextButton(
              child: Text("Annuler"),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
            TextButton(
              child: Text("OK"),
              onPressed: () async {
                try {
                  await TeamService().deletePlayer(widget.eventId, _player.id);
                  widget.onRefresh?.call();
                } on AppException catch (e) {
                  // Handle AppException error
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(content: Text('Erreur: ${e.message}')),
                  );
                } catch (e) {
                  // Handle other errors
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(content: Text('Une erreur est survenue')),
                  );
                }
                context.go('/home', extra: true);
              },
            ),
          ],
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Dialog(
      child: Container(
        padding: const EdgeInsets.all(16),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(
              'Infos joueur',
              style: TextStyle(
                fontSize: 20,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: 8),
            Text(
              _player.name ?? _player.email,
              style: TextStyle(fontSize: 16),
            ),
            const SizedBox(height: 16),
            if (_player.userID != null) ...[
              if (nbEvents != 0) ...[
                Text('Performances pour le ${event?.sport.name.name}',
                    maxLines: 1,
                    overflow: TextOverflow.ellipsis,
                    style: TextStyle(
                      fontSize: 18,
                      fontWeight: FontWeight.bold,
                    )),
                const SizedBox(height: 16),
              ],
              if (nbEvents != 0) ...[
                Text(
                  'Nombre d\'événements : $nbEvents',
                  style: const TextStyle(fontSize: 16),
                ),
                const SizedBox(height: 16)
              ],
              if (stats.isNotEmpty)
                ListView.builder(
                  shrinkWrap: true,
                  itemCount: stats.length,
                  itemBuilder: (context, index) {
                    final stat = stats[index];
                    return ListTile(
                      dense: true,
                      title: Text(stat['stat_label']['label']),
                      trailing: Text(stat['value'].toString()),
                    );
                  },
                )
              else
                const Text('Aucune performance disponible.',
                    style: TextStyle(fontSize: 16)),
              const SizedBox(height: 16)
            ],
            if (widget.canEdit)
              TextButton(
                onPressed: _deletePlayer,
                style: TextButton.styleFrom(
                  foregroundColor: Colors.red.shade300,
                  textStyle: const TextStyle(
                    fontSize: 16,
                    fontWeight: FontWeight.bold,
                    decoration: TextDecoration.underline,
                  ),
                ),
                child: const Text('Supprimer le joueur'),
              ),
            if (!widget.canEdit && widget.isCurrentUser)
              TextButton(
                onPressed: _leaveTeam,
                style: TextButton.styleFrom(
                  foregroundColor: Colors.red.shade300,
                  textStyle: const TextStyle(
                    fontSize: 16,
                    fontWeight: FontWeight.bold,
                    decoration: TextDecoration.underline,
                  ),
                ),
                child: const Text('Quitter l\'équipe (et l\'évent)'),
              ),
          ],
        ),
      ),
    );
  }
}
