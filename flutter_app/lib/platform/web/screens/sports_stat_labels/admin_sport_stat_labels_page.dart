import 'dart:ui';

import 'package:flutter/material.dart';
import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/core/services/sport_stat_labels_service.dart';
import 'package:squad_go/core/services/sport_service.dart';
import 'package:squad_go/core/models/sport_stat_labels.dart';
import 'package:squad_go/platform/web/screens/custom_data_table.dart';
import 'package:squad_go/platform/web/screens/sports_stat_labels/add_edit_sport_stats_labels.dart';

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

  void _confirmDelete(BuildContext context, SportStatLabels statLabel) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return Dialog(
          backgroundColor: Colors
              .transparent, // Rendre le fond transparent pour voir le flou
          child: Stack(
            children: [
              // Appliquer un flou sur l'arrière-plan
              Positioned.fill(
                child: BackdropFilter(
                  filter: ImageFilter.blur(
                      sigmaX: 10.0, sigmaY: 10.0), // Valeur de flou
                  child: Container(
                    color: Colors.black.withOpacity(0), // Pour appliquer un fond transparent
                  ),
                ),
              ),
              // Modal principale
              AlertDialog(
                title: const Text('Supprimer la statistique'),
                content: Text(
                    'Voulez-vous vraiment supprimer la statistique "${statLabel.label}" ?'),
                actions: [
                  TextButton(
                    onPressed: () => Navigator.of(context).pop(),
                    child: const Text(
                      'Annuler',
                      style: TextStyle(color: Colors.black),
                    ),
                  ),
                  TextButton(
                    onPressed: () async {
                      try {
                        await _statLabelsService.deleteStatLabel(statLabel.id);
                        ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(
                              content: Text('La statistique a été supprimée.')),
                        );
                        Navigator.of(context).pop();
                        fetchStatLabels();
                      } catch (e) {
                        ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(
                              content:
                                  Text('Erreur lors de la suppression: $e')),
                        );
                      }
                    },
                    style: TextButton.styleFrom(
                      side: BorderSide(color: Colors.red, width: 2),
                      backgroundColor: Colors.red,
                      padding:
                          EdgeInsets.symmetric(vertical: 12, horizontal: 24),
                    ),
                    child: const Text(
                      'Confirmer',
                      style: TextStyle(color: Colors.white),
                    ),
                  ),
                ],
              ),
            ],
          ),
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _isLoading
          ? const Center(child: CircularProgressIndicator())
          : Column(
              children: [
                Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: Container(
                    padding: const EdgeInsets.symmetric(horizontal: 12.0),
                    decoration: BoxDecoration(
                      color: Colors.blue[50],
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
                        const DropdownMenuItem<String>(
                          value: null,
                          child: Text('Tous les sports'),
                        ),
                        ..._sports.map((sport) => DropdownMenuItem<String>(
                              value: sport['id'].toString(),
                              child: Row(
                                children: [
                                  if (sportIcon.containsKey(SportName.values
                                      .firstWhere((e) =>
                                          e
                                              .toString()
                                              .split('.')
                                              .last
                                              .toLowerCase() ==
                                          sport['name'].toLowerCase())))
                                    Padding(
                                      padding:
                                          const EdgeInsets.only(right: 8.0),
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
                      ],
                      onChanged: (value) {
                        setState(() {
                          _selectedSportId = value;
                          if (value == null) {
                            fetchStatLabels();
                          } else {
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
                Expanded(
                  child: _statLabels.isEmpty
                      ? const Center(child: Text('Aucune donnée disponible'))
                      : CustomDataTable(
                          title: 'Statistiques Sportives',
                          columns: [
                            DataColumn(label: Text('Nom')),
                            DataColumn(label: Text('Unité')),
                            DataColumn(label: Text('Décisif')),
                            DataColumn(label: Text('Actions')),
                          ],
                          rows: _statLabels.map((stat) {
                            return DataRow(cells: [
                              DataCell(Text(stat.label)),
                              DataCell(Text(stat.unit!)),
                              DataCell(Text(stat.isMain! ? 'Oui' : 'Non')),
                              DataCell(
                                Row(
                                  children: [
                                    IconButton(
                                      icon: const Icon(Icons.edit,
                                          color: Colors.blue),
                                      onPressed: () {
                                        showDialog(
                                          context: context,
                                          builder: (context) {
                                            return AddEditSportStatLabelModal(
                                              statLabel:
                                                  stat, // Passe la statistique existante
                                              sports: _sports,
                                              onStatLabelSaved:
                                                  fetchStatLabels, // Rafraîchit après modification
                                            );
                                          },
                                        );
                                      },
                                    ),
                                    IconButton(
                                      icon: const Icon(Icons.delete,
                                          color: Colors.red),
                                      onPressed: () {
                                        _confirmDelete(context, stat);
                                      },
                                    ),
                                  ],
                                ),
                              ),
                            ]);
                          }).toList(),
                          buttonText: 'Ajouter une Statistique',
                          onButtonPressed: () {
                            showDialog(
                              context: context,
                              builder: (context) {
                                return AddEditSportStatLabelModal(
                                  sports: _sports,
                                  onStatLabelSaved: fetchStatLabels,
                                );
                              },
                            );
                          },
                        ),
                )
              ],
            ),
    );
  }
}
