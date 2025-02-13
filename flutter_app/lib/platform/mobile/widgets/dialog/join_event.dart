import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/core/services/team_service.dart';
import 'package:squad_go/main.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';


class JoinEventDialog extends StatefulWidget {
  final String eventId;
  final Future<void> Function()? onRefresh;

  const JoinEventDialog({
    super.key,
    required this.eventId,
    this.onRefresh,
  });

  @override
  State<JoinEventDialog> createState() => _JoinEventDialogState();
}

class _JoinEventDialogState extends State<JoinEventDialog> {
  final EventService eventService = EventService();
  final TeamService teamService = TeamService();
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

  void _joinEvent() async {
    if (event == null || _selectedTeamId == null) return;

    try {
      await teamService.joinTeam(event!.id!, _selectedTeamId!);
      widget.onRefresh?.call();
      context.go('/event/${event!.id}', extra: event);
      Navigator.of(context).pop();
    } catch (e) {
      // Handle error
      log.severe('Failed to join event: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);
    return Dialog(
      child: Container(
        padding: const EdgeInsets.all(16),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(
              '${translate?.join_button ?? "Rejoindre"}: ${event?.name ?? ""}',
              maxLines: 1,
              overflow: TextOverflow.ellipsis,
              style: TextStyle(
                fontSize: 20,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: 16),
            Text(translate?.choose_team ?? 'Choisissez votre équipe:'),
            const SizedBox(height: 16),
            SizedBox(
              height: 75 * MediaQuery.of(context).devicePixelRatio,
              child: SingleChildScrollView(
                child: Column(
                  children: event?.teams
                          ?.map((team) => RadioListTile<String>(
                                title: Text(team.name),
                                subtitle: Text(
                                  '${team.players.length}${team.maxPlayers > 0 ? '/${team.maxPlayers}' : ''} ${translate?.sporty ?? "sportifs"}',                                ),
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
              ),
            ),
            const SizedBox(height: 16),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                ElevatedButton(
                  onPressed: () {
                    Navigator.of(context).pop();
                  },
                  style: ElevatedButton.styleFrom(
                    padding: const EdgeInsets.symmetric(
                        horizontal: 24, vertical: 12),
                  ),
                  child: Text(translate?.cancel ?? 'Annuler'),
                ),
                const SizedBox(width: 16),
                ElevatedButton(
                  onPressed: _selectedTeamId == null
                      ? null
                      : () {
                          // Join the event
                          _joinEvent();
                        },
                  style: ElevatedButton.styleFrom(
                    padding: const EdgeInsets.symmetric(
                        horizontal: 24, vertical: 12),
                  ),
                  child: Text(translate?.join_button ?? 'Rejoindre'),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
