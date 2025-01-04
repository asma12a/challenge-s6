import 'dart:convert';
import 'dart:developer';

import 'package:flutter/cupertino.dart';
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

  @override
  void initState() {
    super.initState();
    _fetchEventDetails();
    _fetchStatLabels();
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

  Future<void> _fetchStatLabels() async {
    try {
      final labels = await statLabelsService.getStatLabelsBySport(widget.event.sport.id);
      setState(() {
        statLabels = labels;
        ratings = labels.map((label) {
          return UserStats(id: label.id, value: 0, stat: label);
        }).toList();
      });
    } catch (e) {}
  }

  Map<String, dynamic> _generateJson() {
    List<Map<String, dynamic>> stats = ratings.where((rating) => rating.value > 0).map((rating) {
      return {
        'stat_id': rating.stat?.id,
        'stat_value': rating.value,
      };
    }).toList();

    return {
      'user_id': '01JCBH35S8EDHF9FYKER9YJ075',
      'stats': stats,
    };
  }

  void _addUserStat() async {
    try{
      await statLabelsService.addUserStat(_generateJson(), widget.event.id!);
      ScaffoldMessenger.of(context).clearSnackBars();
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text(
            textAlign: TextAlign.center,
            "L'événement a bien été enregistré.",
            style: TextStyle(color: Theme.of(context).colorScheme.onPrimary),
          ),
          backgroundColor: Theme.of(context).colorScheme.primary,
        ),
      );
    }catch (e) {}
  }



  @override
  Widget build(BuildContext context) {
    return Dialog(
      child: Container(
        padding: const EdgeInsets.all(16),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Text('${event?.name ?? 'l\'événement'}',
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                )),
            const SizedBox(height: 16),
            Text('Noter les performances de player'),
            const SizedBox(height: 16),
            Expanded(
              child: ListView.builder(
                shrinkWrap: true,
                itemCount: statLabels.length,
                itemBuilder: (context, index) {
                  final statLabel = statLabels[index];


                  // Crée un UserStats pour chaque label avec valeur initiale de 0
                  final userStat = ratings.firstWhere(
                    (stat) => stat.stat?.id == statLabel.id,
                    orElse: () => UserStats(id: statLabel.id, value: 0, stat: statLabel),
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
                    final jsonData = _generateJson();
                    debugPrint(jsonEncode(jsonData));
                    _addUserStat();
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
