import 'package:flutter/material.dart';
import 'package:squad_go/core/services/sport_stat_labels_service.dart';
import 'package:squad_go/core/services/sport_service.dart';
import 'package:squad_go/core/models/sport_stat_labels.dart';
import '../custom_data_table.dart';

class AdminSportStatLabelsPage extends StatefulWidget {
  const AdminSportStatLabelsPage({super.key});

  @override
  State<AdminSportStatLabelsPage> createState() =>
      _AdminSportStatLabelsPageState();
}

class _AdminSportStatLabelsPageState extends State<AdminSportStatLabelsPage> {
  final SportStatLabelsService _statLabelsService = SportStatLabelsService();
  List<SportStatLabels> _statLabels = [];
  List<Map<String, dynamic>> _sports = [];
  String? _selectedSportId;
  bool _isLoading = true;

  @override
  void initState() {
    super.initState();
    fetchSports();
  }

  /// Récupérer la liste des sports
  Future<void> fetchSports() async {
    try {
      setState(() {
        _isLoading = true;
      });
      final sports = await SportService.getSports();
      setState(() {
        _sports = sports;
        _isLoading = false;
      });
    } catch (e) {
      setState(() {
        _isLoading = false;
      });
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Erreur lors du chargement des sports: $e')),
      );
    }
  }

  /// Récupérer les statistiques pour un sport donné
  Future<void> fetchStatLabels(String sportId) async {
    try {
      setState(() {
        _isLoading = true;
      });
      final labels = await _statLabelsService.getStatLabelsBySport(sportId);
      setState(() {
        _statLabels = labels;
        _isLoading = false;
      });
    } catch (e) {
      setState(() {
        _isLoading = false;
      });
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Erreur lors du chargement des stats: $e')),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
    
      body: _isLoading
          ? const Center(child: CircularProgressIndicator())
          : Column(
              children: [
                // Dropdown pour sélectionner un sport
                Padding(
                  padding: const EdgeInsets.all(8.0),
                  child: DropdownButtonFormField<String>(
                    value: _selectedSportId,
                    hint: const Text('Sélectionnez un sport'),
                    items: _sports
                        .map((sport) => DropdownMenuItem<String>(
                              value: sport['id'].toString(),
                              child: Text(sport['name']),
                            ))
                        .toList(),
                    onChanged: (value) {
                      if (value != null) {
                        setState(() {
                          _selectedSportId = value;
                          _statLabels = [];
                        });
                        fetchStatLabels(value);
                      }
                    },
                  ),
                ),
                // Utilisation de CustomDataTable
                Expanded(
                  child: _statLabels.isEmpty
                      ? const Center(child: Text('Aucune donnée disponible'))
                      : CustomDataTable(
                          title: 'Statistiques sportives',
                          columns: [
                            DataColumn(label: Text('Nom')),
                            DataColumn(label: Text('Unité')),
                            DataColumn(label: Text('Décisif')),
                          ],
                          rows: _statLabels.map((stat) {
                            return DataRow(cells: [
                              DataCell(Text(stat.label)),
                              DataCell(Text(stat.unit)),
                              DataCell(Text(stat.isMain ? 'Oui' : 'Non')),
                            ]);
                          }).toList(),
                          buttonText: 'Ajouter une Statistique',
                          onButtonPressed: () {
                            // Ajouter la logique pour ajouter une nouvelle statistique
                          },
                        ),
                ),
              ],
            ),
    );
  }
}
