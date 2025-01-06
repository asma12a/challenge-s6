import 'package:flutter/material.dart';
import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/core/services/sport_stat_labels_service.dart';
import 'package:squad_go/core/services/sport_service.dart';
import 'package:squad_go/core/models/sport_stat_labels.dart';
import 'package:squad_go/platform/web/screens/custom_data_table.dart';

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
    fetchStatLabels();
  }

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

  Future<void> fetchStatLabels() async {
    try {
      setState(() {
        _isLoading = true;
      });
      final labels = await _statLabelsService.getAllStatLabels();
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

  Future<void> fetchStatLabelsBySport(String sportId) async {
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
        SnackBar(
            content:
                Text('Erreur lors du chargement des stats pour ce sport: $e')),
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
                // Container amélioré pour le dropdown
                Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: Container(
                    padding: const EdgeInsets.symmetric(horizontal: 12.0),
                    decoration: BoxDecoration(
                      color: Colors.blue[50], // Arrière-plan léger
                      borderRadius: BorderRadius.circular(12.0),
                      boxShadow: [
                        BoxShadow(
                          color: Colors.black26,
                          blurRadius: 8.0,
                          offset: Offset(2, 2),
                        ),
                      ],
                    ),
                    child: DropdownButtonFormField<String>(
                      value: _selectedSportId,
                      hint: const Text(
                        'Sélectionnez un sport',
                        style: TextStyle(color: Colors.blueGrey),
                      ),
                      items: [
                        // Option "Tous les sports"
                        const DropdownMenuItem<String>(
                          value: null, // null signifie "Tous les sports"
                          child: Text('Tous les sports'),
                        ),
                        // Liste des sports disponibles
                        ..._sports
                            .map((sport) => DropdownMenuItem<String>(
                                  value: sport['id'].toString(),
                                  child: Row(
                                    children: [
                                      // Vérifie si l'icône existe dans le Map sportIcon
                                      if (sportIcon.containsKey(SportName.values
                                          .firstWhere((e) =>
                                              e
                                                  .toString()
                                                  .split('.')
                                                  .last
                                                  .toLowerCase() ==
                                              sport['name'].toLowerCase())))
                                        Padding(
                                          padding: const EdgeInsets.only(right: 8.0),
                                          child: Icon(
                                            sportIcon[SportName.values.firstWhere(
                                                (e) =>
                                                    e
                                                        .toString()
                                                        .split('.')
                                                        .last
                                                        .toLowerCase() ==
                                                    sport['name'].toLowerCase())],
                                            size: 20,
                                            color: Colors.blue,
                                          ),
                                        ),
                                      Text(
                                        sport['name'],
                                        style: const TextStyle(fontSize: 16),
                                      ),
                                    ],
                                  ),
                                ))
                            .toList(),
                      ],
                      onChanged: (value) {
                        setState(() {
                          _selectedSportId = value; // La valeur peut être null
                          if (value == null) {
                            // Si "Tous les sports" est sélectionné
                            fetchStatLabels(); // Charger toutes les statistiques
                          } else {
                            // Sinon, charger les statistiques pour un sport spécifique
                            fetchStatLabelsBySport(value);
                          }
                        });
                      },
                      decoration: InputDecoration(
                        border: InputBorder.none,
                        icon: const Icon(Icons.sports, color: Colors.blue),
                      ),
                    ),
                  ),
                ),
                // Utilisation de CustomDataTable
                Expanded(
                  child: _statLabels.isEmpty
                      ? const Center(child: Text('Aucune donnée disponible'))
                      : CustomDataTable(
                          title: 'Statistiques Sportives',
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
