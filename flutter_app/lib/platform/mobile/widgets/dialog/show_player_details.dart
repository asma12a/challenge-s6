import 'package:flutter/material.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/services/sport_stat_labels_service.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/models/user_stats.dart';

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
  List<UserStats> ratings = [];

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
    } catch (e) {
      // Handle error
      log.severe('Failed to fetch event details: $e');
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
            Text('Performances pour le ${event?.sport.name.name}',
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
                style: TextStyle(
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                )),
          ],
        ),
      ),
    );
  }
}
