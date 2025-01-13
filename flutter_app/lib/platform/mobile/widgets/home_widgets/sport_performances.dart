import 'package:flutter/material.dart';
import 'package:squad_go/core/services/sport_stat_labels_service.dart';
import 'package:squad_go/main.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';


class SportPerformances extends StatefulWidget {
  final String sportId;
  final String userId;

  const SportPerformances(
      {super.key, required this.sportId, required this.userId});

  @override
  State<SportPerformances> createState() => _SportPerformancesState();
}

class _SportPerformancesState extends State<SportPerformances> {
  final SportStatLabelsService statLabelsService = SportStatLabelsService();

  int nbEvents = 0;
  List<Map<String, dynamic>> stats = [];

  @override
  void initState() {
    super.initState();
    _loadUserPerformance();
  }

  Future<void> _loadUserPerformance() async {
    try {
      final userPerformances = await statLabelsService
          .getUserPerformanceBySport(widget.sportId, widget.userId);
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
    final translate = AppLocalizations.of(context);

    return Container(
        margin: const EdgeInsets.only(top: 16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            if (nbEvents != 0) ...[
              Text(
                '${translate?.nb_events ?? "Nombre d\'événements :"} $nbEvents',
                style: const TextStyle(fontSize: 16),
              ),
              const SizedBox(height: 16),
            ],
            if (stats.isNotEmpty)
              ListView.builder(
                shrinkWrap: true,
                itemCount: stats.length,
                itemBuilder: (context, index) {
                  final stat = stats[index];
                  return Container(
                    margin: EdgeInsets.only(bottom: 16),
                    decoration: BoxDecoration(
                      color: Theme.of(context)
                          .colorScheme
                          .primary
                          .withAlpha(20),
                      borderRadius: BorderRadius.circular(16),
                    ),
                    child: ListTile(
                      dense: true,
                      title: Text(stat['stat_label']['label']),
                      trailing: Text(stat['value'].toString()),
                    ),
                  );
                },
              )
            else
              Text(translate?.no_perf ?? 'Aucune performance disponible.',
                  style: const TextStyle(fontSize: 16)),
          ],
        ));
  }
}
