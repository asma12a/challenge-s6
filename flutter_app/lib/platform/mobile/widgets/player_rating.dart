import 'package:flutter/material.dart';
import 'package:squad_go/core/models/user_stats.dart';
import 'package:squad_go/core/services/sport_stat_labels_service.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/models/sport_stat_labels.dart';

class PlayerRating extends StatefulWidget {
  final Event event;

  const PlayerRating({super.key, required this.event});

  @override
  State<PlayerRating> createState() => _PlayerRatingState();
}

class _PlayerRatingState extends State<PlayerRating> {
  final SportStatLabelsService statLabelsService = SportStatLabelsService();
  Event? event;
  List<SportStatLabels> statLabels = [];
  List<UserStats> ratings = [];
  bool isLoading = true;
  bool isUpdating = false;

  @override
  void initState() {
    super.initState();
    _fetchEventDetails();
    _initializeRatings();
  }

  void _fetchEventDetails() async {
    try {
      setState(() {
        event = widget.event;
      });
    } catch (e) {
      // Handle error
    }
  }

  Future<void> _initializeRatings() async {
    try {
      setState(() => isLoading = true);

      // Charger les labels
      final labels = await statLabelsService.getStatLabelsBySport(widget.event.sport.id);
      setState(() {
        statLabels = labels;
      });

      // Charger les UserStats existants pour l'utilisateur
      await _loadEventRatings();

      // Si aucun UserStat existant, créer des entrées par défaut
      if (ratings.isEmpty) {
        setState(() {
          ratings = labels.map((label) {
            return UserStats(id: label.id, value: 0, stat: label);
          }).toList();
        });
      }else{
        isUpdating = true;
      }
    } catch (e) {
      // Gérer l'erreur
    } finally {
      setState(() => isLoading = false);
    }
  }


  Future<void> _loadEventRatings() async {
    try {
      if (widget.event.id == null) return;
      final existingRatings =
          await statLabelsService.getUserStatByEvent(widget.event.id!, '01JCBH35S8EDHF9FYKER9YJ075');
      setState(() {
        ratings = existingRatings;
      });
    } catch (e) {}
  }


  void _saveUserStat() async {
    try {

      if(isUpdating){
        final stats = ratings.map((rating) {
          return {
            'user_stat_id': rating.id,
            'stat_value': rating.value,
          };
        }).toList();

        final jsonData = {
          'stats': stats,
        };
        await statLabelsService.updateUserStat(jsonData);

      }else{
        final stats = ratings.map((rating) {
          return {
            'stat_id': rating.stat?.id,
            'stat_value': rating.value,
          };
        }).toList();

        final jsonData = {
          'user_id': '01JCBH35S8EDHF9FYKER9YJ075',
          'stats': stats,
        };
        await statLabelsService.addUserStat(jsonData, widget.event.id!);
      }

      Navigator.of(context).pop();
    } catch (e) {}
  }

  @override
  Widget build(BuildContext context) {
    return Dialog(
      child: Container(
        padding: const EdgeInsets.all(16),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(event?.name ?? 'l\'événement',
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                )),
            const SizedBox(height: 16),
            Text('Noter les performances du player'),
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
                    orElse: () => UserStats(id: null, value: 0, stat: statLabel),
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
                                  final index = ratings.indexWhere((stat) => stat.stat?.id == statLabel.id);
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
                    padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 12),
                  ),
                  child: const Text('Annuler'),
                ),
                const SizedBox(width: 16),
                ElevatedButton(
                  onPressed: () {
                    _saveUserStat();
                  },
                  style: ElevatedButton.styleFrom(
                    padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 12),
                  ),
                  child: const Text('Soumettre'),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
