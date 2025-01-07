import 'package:flutter/material.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/services/sport_stat_labels_service.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/core/models/event.dart';

import '../../../../main.dart';

// TODO: Full details of player: player info + player stats
// TODO: If coach or org : can delete player from team
class ShowPlayerDetailsDialog extends StatefulWidget {
  final String eventId;
  final Player player;
  final Future<void> Function()? onRefresh;

  const ShowPlayerDetailsDialog({super.key, required this.eventId, required this.player, this.onRefresh});

  @override
  State<ShowPlayerDetailsDialog> createState() => _ShowPlayerDetailsDialogState();
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

      final userPerformances = await statLabelsService.getUserPerformanceBySport(event!.sport.id, _player.userID);
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

  @override
  Widget build(BuildContext context) {
    return Dialog(
      child: Container(
        padding: const EdgeInsets.all(16),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(
              'Nom : ${_player.name}',
              style: TextStyle(
                fontSize: 20,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: 16),
            if (nbEvents != 0)
              Text('Performances pour le ${event?.sport.name.name}',
                  maxLines: 1,
                  overflow: TextOverflow.ellipsis,
                  style: TextStyle(
                    fontSize: 18,
                    fontWeight: FontWeight.bold,
                  )),
            const SizedBox(height: 16),
            if (nbEvents != 0)
              Text(
                'Nombre d\'événements : $nbEvents',
                style: const TextStyle(fontSize: 16),
              ),
            const SizedBox(height: 16),
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
              const Text('Aucune performance disponible.', style: TextStyle(fontSize: 16)),
          ],
        ),
      ),
    );
  }
}
