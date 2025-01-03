import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:squad_go/core/models/user_stats.dart';
import 'package:squad_go/core/services/sport_stat_labels_service.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/models/sport_stat_labels.dart';

class PlayerRating extends StatefulWidget {
  final Event event;

  const PlayerRating({
    super.key,
    required this.event
  });

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
    _loadEventRatings();
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

  void _fetchStatLabels() async {
    try {
      final labels = await SportStatLabelsService.getStatLabelsBySport(widget.event.sport.id);
      setState(() {
        statLabels = labels;
      });

    } catch (e) {
    }
  }

  void _loadEventRatings() async {
    try {
      if (widget.event.id == null) return;
      final labels = await SportStatLabelsService.getUserStatByEvent(widget.event.id!, '01JCBH35S8EDHF9FYKER9YJ075');
      setState(() {
        ratings = labels;
      });

    } catch (e) {
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
            Text('Noter ${event?.name ?? 'l\'événement'}',
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                )),
            const SizedBox(height: 16),
            Text('Noter les performances de player'),
            const SizedBox(height: 16),
            SizedBox(
              height: 75 * MediaQuery.of(context).devicePixelRatio,
              child: SingleChildScrollView(
                child: Column(
                  children: statLabels.map((statLabel) {
                    return Padding(
                      padding: const EdgeInsets.symmetric(vertical: 8.0),
                      child: Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          Text(statLabel.label),
                          Row(
                            children: [
                              IconButton(
                                icon: Icon(Icons.remove),
                                onPressed: () {

                                },
                              ),

                            ],
                          ),
                        ],
                      ),
                    );
                  }).toList(),
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
                    padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 12),
                  ),
                  child: const Text('Annuler'),
                ),
                const SizedBox(width: 16),
                ElevatedButton(
                  onPressed: null,
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
