import 'dart:ui';
import 'package:flutter/material.dart';
import 'package:squad_go/core/services/event_service.dart';
import '../custom_data_table.dart';
import './add_edit_event_page.dart';

class AdminEventsPage extends StatefulWidget {
  const AdminEventsPage({super.key});

  @override
  State<AdminEventsPage> createState() => _AdminEventsPageState();
}

class _AdminEventsPageState extends State<AdminEventsPage> {
  List<Map<String, dynamic>> events = [];
  bool isLoading = true;

  @override
  void initState() {
    super.initState();
    fetchEvents();
  }

  Future<void> fetchEvents() async {
    try {
      final fetchedEvents = await EventService.getEvents();
      setState(() {
        events = fetchedEvents;
        isLoading = false;
      });
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Erreur: $e')),
      );
    }
  }

  Future<void> deleteEvent(String id) async {
    try {
      await EventService.deleteEvent(id);
      if (mounted) {
        fetchEvents();
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Événement supprimé avec succès.')),
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

  Future<void> confirmDelete(String eventName, String id) async {
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
                title: const Text('Supprimer l\'événement'),
                content: Text('Voulez-vous vraiment supprimer "$eventName" ?'),
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
                      // Effectuer la suppression avant de fermer le dialogue
                      await deleteEvent(id);
                      Navigator.of(context).pop(); // Fermer le dialogue
                      if (mounted) {
                        ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(content: Text('$eventName a été supprimé.')),
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
              title: 'Événements',
              columns: const [
                DataColumn(label: Text('Nom')),
                DataColumn(label: Text('Date')),
                DataColumn(label: Text('Statut')),
                DataColumn(label: Text('Code')),
                DataColumn(label: Text('Terminé')),
                DataColumn(label: Text('Adresse')),
                DataColumn(label: Text('Actions')),
              ],
              rows: events.map((event) {
                return DataRow(cells: [
                  DataCell(Text(event['name'])),
                  DataCell(Text(event['date'])),
                  DataCell(
                      Text(event['is_public'] ?? false ? 'Public' : 'Privé')),
                  DataCell(Text(event['event_code'])),
                  DataCell(Text(event['is_finished'] ?? false ? 'Oui' : 'Non')),
                  DataCell(Text(event['address'] ?? 'N/A')),
                  DataCell(Row(
                    mainAxisAlignment: MainAxisAlignment.end,
                    children: [
                      IconButton(
                        icon: const Icon(Icons.edit, color: Colors.blue),
                        onPressed: () {
                          showDialog(
                            context: context,
                            builder: (context) {
                              return AddEditEventModal(
                                event: event,
                                onEventSaved: fetchEvents,
                              );
                            },
                          );
                        },
                      ),
                      IconButton(
                        icon: const Icon(Icons.delete, color: Colors.red),
                        onPressed: () =>
                            confirmDelete(event['name'], event['id']),
                      ),
                    ],
                  )),
                ]);
              }).toList(),
              buttonText: 'Ajouter un événement',
              onButtonPressed: () {
                showDialog(
                  context: context,
                  builder: (context) {
                    return AddEditEventModal(
                      event: null,
                      onEventSaved: fetchEvents,
                    );
                  },
                );
              },
            ),
    );
  }
}
