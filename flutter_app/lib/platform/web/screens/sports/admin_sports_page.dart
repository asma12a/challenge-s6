import 'dart:ui';
import 'package:flutter/material.dart';
import 'package:squad_go/core/services/sport_service.dart';
import '../custom_data_table.dart';
import './add_edit_sport_page.dart';

class AdminSportsPage extends StatefulWidget {
  const AdminSportsPage({super.key});

  @override
  State<AdminSportsPage> createState() => _AdminSportsPageState();
}

class _AdminSportsPageState extends State<AdminSportsPage> {
  List<Map<String, dynamic>> sports = [];
  bool isLoading = true;

  @override
  void initState() {
    super.initState();
    fetchSports();
  }

  Future<void> fetchSports() async {
    try {
      final fetchedSports = await SportService.getSports();
      setState(() {
        sports = fetchedSports;
        isLoading = false;
      });
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Erreur: $e')),
      );
    }
  }

  Future<void> deleteSport(String id) async {
    try {
      await SportService.deleteSport(id);
      if (mounted) {
        fetchSports();
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Sport supprimé avec succès.')),
        );
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Erreur: $e')),
        );
      }
    }
  }

  Future<void> confirmDelete(String sportName, String id) async {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return Dialog(
          backgroundColor: Colors.transparent,
          child: Stack(
            children: [
              Positioned.fill(
                child: BackdropFilter(
                  filter: ImageFilter.blur(sigmaX: 10.0, sigmaY: 10.0),
                  child: Container(
                    color: Colors.black.withOpacity(0),
                  ),
                ),
              ),
              AlertDialog(
                title: const Text('Supprimer le sport'),
                content: Text('Voulez-vous vraiment supprimer "$sportName" ?'),
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
                      await deleteSport(id);
                      Navigator.of(context).pop();
                      if (mounted) {
                        ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(content: Text('$sportName a été supprimé.')),
                        );
                      }
                    },
                    style: TextButton.styleFrom(
                      side: const BorderSide(color: Colors.red, width: 2),
                      backgroundColor: Colors.red,
                      padding: const EdgeInsets.symmetric(
                          vertical: 12, horizontal: 24),
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
      body: isLoading
          ? const Center(child: CircularProgressIndicator())
          : CustomDataTable(
              title: 'Sports',
              columns: const [
                DataColumn(label: Text('Nom')),
                DataColumn(label: Text('Actions')),
              ],
              rows: sports.map((sport) {
                return DataRow(cells: [
                  DataCell(Text(sport['name'])),
                  DataCell(Row(
                    mainAxisAlignment: MainAxisAlignment.end,
                    children: [
                      IconButton(
                        icon: const Icon(Icons.edit, color: Colors.blue),
                        onPressed: () {
                          showDialog(
                            context: context,
                            builder: (context) {
                              return AddEditSportModal(
                                sport: sport,
                                onSportSaved: fetchSports,
                              );
                            },
                          );
                        },
                      ),
                      IconButton(
                        icon: const Icon(Icons.delete, color: Colors.red),
                        onPressed: () =>
                            confirmDelete(sport['name'], sport['id']),
                      ),
                    ],
                  )),
                ]);
              }).toList(),
              buttonText: 'Ajouter un sport',
              onButtonPressed: () {
                showDialog(
                  context: context,
                  builder: (context) {
                    return AddEditSportModal(
                      sport: null,
                      onSportSaved: fetchSports,
                    );
                  },
                );
              },
            ),
    );
  }
}
