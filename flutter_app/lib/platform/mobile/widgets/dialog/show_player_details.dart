import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/providers/connectivity_provider.dart';
import 'package:squad_go/core/services/sport_stat_labels_service.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/services/team_service.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/offline.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

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
    final translate = AppLocalizations.of(context);
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: Text(translate?.del_player ?? 'Supprimer le joueur'),
          content: Text(
              '${translate?.confirm_del_player ?? "Êtes-vous sûr de vouloir supprimer le joueur"} ${_player.name ?? ""} ?'),
          actions: [
            TextButton(
              child: Text(translate?.cancel ?? "Annuler"),
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
                    SnackBar(
                      content: Text(
                        '${translate?.error ?? "Erreur:"} ${e.message}',
                      ),
                    ),
                  );
                } catch (e) {
                  // Handle other errors
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                        content: Text(translate?.error_occurred ??
                            'Une erreur est survenue')),
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
    final translate = AppLocalizations.of(context);

    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: Text(translate?.leave_team ?? 'Quitter l\'équipe'),
          content: Text(translate?.leave_team_event ??
              'Êtes-vous sûr de vouloir quitter l\'équipe et l\'événement ?'),
          actions: [
            TextButton(
              child: Text(translate?.cancel ?? "Annuler"),
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
                    SnackBar(
                      content: Text(
                        '${translate?.error ?? "Erreur:"} ${e.message}',
                      ),
                    ),
                  );
                } catch (e) {
                  // Handle other errors
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                        content: Text(translate?.error_occurred ??
                            'Une erreur est survenue')),
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
    var isOnline = context.watch<ConnectivityState>().isConnected;
    final translate = AppLocalizations.of(context);
    final Color roleColor = widget.player.role == PlayerRole.coach
        ? Colors.blue
        : widget.player.role == PlayerRole.org ||
                event?.createdBy == _player.userID
            ? Colors.deepOrange
            : Colors.green;
    return Dialog(
      child: Container(
        padding: const EdgeInsets.all(16),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(
              translate?.infos_player ?? 'Infos participant',
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
            const SizedBox(height: 8),
            Badge(
              label: Text(
                widget.player.role == PlayerRole.coach
                    ? 'Coach'
                    : widget.player.role == PlayerRole.org ||
                            event?.createdBy == _player.userID
                        ? translate?.organizer ?? 'Organisateur'
                        : translate?.player ?? 'Joueur',
                style: TextStyle(
                  color: roleColor,
                  fontWeight: FontWeight.bold,
                ),
              ),
              backgroundColor: roleColor.withAlpha(40),
              padding: EdgeInsets.symmetric(
                horizontal: 5,
                vertical: 2,
              ),
            ),
            const SizedBox(height: 16),
            if (_player.userID != null) ...[
              if (nbEvents != 0) ...[
                Text(
                  '${translate?.perf_for ?? "Performances pour le"} ${event?.sport.name.name}',
                  maxLines: 1,
                  overflow: TextOverflow.ellipsis,
                  style: TextStyle(
                    fontSize: 18,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                const SizedBox(height: 16),
              ],
              if (nbEvents != 0) ...[
                Text(
                  '${translate?.nb_events ?? "Nombre d'événements :"} $nbEvents',
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
                Text(translate?.no_perf ?? 'Aucune performance disponible.',
                    style: const TextStyle(fontSize: 16)),
              const SizedBox(height: 16)
            ],
            if (widget.canEdit)
              TextButton(
                onPressed: () {
                  if (!isOnline) {
                    showDialog(
                      context: context,
                      builder: (context) => const OfflineDialog(),
                    );
                    return;
                  }
                  _deletePlayer();
                },
                style: TextButton.styleFrom(
                  foregroundColor: Colors.red.shade300,
                  textStyle: const TextStyle(
                    fontSize: 16,
                    fontWeight: FontWeight.bold,
                    decoration: TextDecoration.underline,
                  ),
                ),
                child: Text(translate?.del_player ?? 'Supprimer le joueur'),
              ),
            if (!widget.canEdit && widget.isCurrentUser)
              TextButton(
                onPressed: () {
                  if (!isOnline) {
                    showDialog(
                      context: context,
                      builder: (context) => const OfflineDialog(),
                    );
                    return;
                  }
                  _leaveTeam();
                },
                style: TextButton.styleFrom(
                  foregroundColor: Colors.red.shade300,
                  textStyle: const TextStyle(
                    fontSize: 16,
                    fontWeight: FontWeight.bold,
                    decoration: TextDecoration.underline,
                  ),
                ),
                child: Text(translate?.leave_team_and_event ??
                    'Quitter l\'équipe (et l\'évent)'),
              ),
          ],
        ),
      ),
    );
  }
}
