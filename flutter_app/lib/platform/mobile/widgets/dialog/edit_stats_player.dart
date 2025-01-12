import 'package:flutter/material.dart';
import 'package:squad_go/core/models/user_stats.dart';
import 'package:squad_go/core/services/sport_stat_labels_service.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/models/sport_stat_labels.dart';

import 'package:squad_go/core/services/event_service.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import '../../../../main.dart';

class EditStatsPlayerDialog extends StatefulWidget {
  final String eventId;
  final Player player;
  final Future<void> Function()? onRefresh;

  const EditStatsPlayerDialog(
      {super.key, required this.eventId, required this.player, this.onRefresh});

  @override
  State<EditStatsPlayerDialog> createState() => _EditStatsPlayerDialogState();
}

class _EditStatsPlayerDialogState extends State<EditStatsPlayerDialog> {
  final SportStatLabelsService statLabelsService = SportStatLabelsService();
  final EventService eventService = EventService();
  Event? event;
  late Player _player;
  List<SportStatLabels> statLabels = [];
  List<UserStats> ratings = [];
  bool isLoading = true;
  bool isUpdating = false;

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
      _initializeRatings();
    } catch (e) {
      // Handle error
      log.severe('Failed to fetch event details: $e');
    }
  }

  Future<void> _initializeRatings() async {
    try {
      if (event == null) return;
      setState(() => isLoading = true);

      final labels =
          await statLabelsService.getStatLabelsBySport(event!.sport.id);
      setState(() {
        statLabels = labels;
      });
      await _loadEventRatings();

      if (ratings.isEmpty) {
        setState(() {
          ratings = labels.map((label) {
            return UserStats(id: label.id, value: 0, stat: label);
          }).toList();
        });
      } else {
        isUpdating = true;
      }
    } catch (e) {
      log.severe('Failed to fetch stat labels details: $e');
    } finally {
      setState(() => isLoading = false);
    }
  }

  Future<void> _loadEventRatings() async {
    try {
      if (event!.id == null) return;
      final existingRatings = await statLabelsService.getUserStatByEvent(
          event!.id!, _player.userID);
      setState(() {
        ratings = existingRatings;
      });
    } catch (e) {
      log.severe('Failed to fetch user stat: $e');
    }
  }

  void _saveUserStat() async {
    final translate = AppLocalizations.of(context);
    try {
      if (isUpdating) {
        final stats = ratings.map((rating) {
          return {
            'user_stat_id': rating.id,
            'stat_value': rating.value,
          };
        }).toList();

        final jsonData = {
          'stats': stats,
        };
        await statLabelsService.updateUserStat(jsonData, event!.id!);
      } else {
        final stats = ratings.map((rating) {
          return {
            'stat_id': rating.stat?.id,
            'stat_value': rating.value,
          };
        }).toList();

        final jsonData = {
          'user_id': _player.userID!,
          'stats': stats,
        };
        await statLabelsService.addUserStat(jsonData, event!.id!);
      }
      widget.onRefresh?.call();
      Navigator.of(context).pop();
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
            content:
                Text(translate?.note_success ?? 'Ajout des notes avec succès')),
      );
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
            content:
                Text(translate?.error_occurred ?? 'Une erreur est survenue')),
      );
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
              event?.name ?? translate!.event,
              maxLines: 1,
              overflow: TextOverflow.ellipsis,
              style: TextStyle(
                fontSize: 20,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: 16),
            Text(
              '${translate?.note_player ?? "Noter les performances du joueur"} ${_player.name ?? ""}',
              maxLines: 1,
              overflow: TextOverflow.ellipsis,
              style: TextStyle(
                fontSize: 20,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: 16),
            ConstrainedBox(
              constraints: BoxConstraints(
                maxHeight: 300,
              ),
              child: ListView.builder(
                shrinkWrap: true,
                itemCount: statLabels.length,
                itemBuilder: (context, index) {
                  final statLabel = statLabels[index];

                  final userStat = ratings.firstWhere(
                    (stat) => stat.stat?.id == statLabel.id,
                    orElse: () =>
                        UserStats(id: null, value: 0, stat: statLabel),
                  );

                  return Padding(
                    padding: const EdgeInsets.symmetric(vertical: 8.0),
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Text(
                          statLabel.label,
                          style: TextStyle(fontSize: 16),
                        ),
                        Row(
                          children: [
                            IconButton(
                              icon: Icon(Icons.remove),
                              onPressed: () {
                                setState(() {
                                  if (userStat.value > 0) {
                                    userStat.value--;
                                  }
                                });
                              },
                            ),
                            Text(
                              userStat.value.toString(),
                              style: TextStyle(fontSize: 16),
                            ),
                            IconButton(
                              icon: Icon(Icons.add),
                              onPressed: () {
                                setState(() {
                                  final index = ratings.indexWhere(
                                      (stat) => stat.stat?.id == statLabel.id);
                                  if (index != -1) {
                                    ratings[index].value++;
                                  }
                                });
                              },
                            ),
                          ],
                        ),
                      ],
                    ),
                  );
                },
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
                  onPressed: () {
                    _saveUserStat();
                  },
                  style: ElevatedButton.styleFrom(
                    foregroundColor: Colors.white,
                    backgroundColor: Colors.blue,
                    padding: const EdgeInsets.symmetric(
                        horizontal: 24, vertical: 12),
                  ),
                  child: Text(translate?.save_event ?? 'Soumettre'),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
