import 'package:flutter/material.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/main.dart';

class JoinEventDialog extends StatefulWidget {
  final String eventId;
  const JoinEventDialog({super.key, required this.eventId});

  @override
  State<JoinEventDialog> createState() => _JoinEventDialogState();
}

class _JoinEventDialogState extends State<JoinEventDialog> {
  final EventService eventService = EventService();
  Event? event;
  String? _selectedTeamId;

  @override
  void initState() {
    super.initState();

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
            Text('Rejoindre: ${event?.name}',
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                )),
            const SizedBox(height: 16),
            Text('Choisissez votre équipe:'),
            const SizedBox(height: 16),
            Column(
              children: event?.teams
                      ?.map((team) => RadioListTile<String>(
                            title: Text(team.name),
                            value: team.id,
                            groupValue: _selectedTeamId,
                            onChanged: (String? newValue) {
                              setState(() {
                                _selectedTeamId = newValue;
                              });
                            },
                          ))
                      .toList() ??
                  [],
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                TextButton(
                  onPressed: () {
                    Navigator.of(context).pop();
                  },
                  child: const Text('Annuler'),
                ),
                TextButton(
                  onPressed: () {
                    // Join the event
                    Navigator.of(context).pop();
                  },
                  child: const Text('Rejoindre'),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
