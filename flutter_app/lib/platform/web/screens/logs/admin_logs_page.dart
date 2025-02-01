import 'package:flutter/material.dart';
import 'package:squad_go/core/services/log_service.dart';
import 'package:intl/intl.dart';
import 'package:squad_go/platform/web/screens/custom_data_table.dart';

class AdminLogsPage extends StatefulWidget {
  const AdminLogsPage({super.key});

  @override
  State<AdminLogsPage> createState() => _AdminLogsPageState();
}

class _AdminLogsPageState extends State<AdminLogsPage> {
  List<Map<String, dynamic>> logs = [];
  bool isLoading = true;

  @override
  void initState() {
    super.initState();
    fetchLogs();
  }

  Future<void> fetchLogs() async {
    try {
      final fetchedLogs = await LogService.getLogs();
      if (mounted) {
        setState(() {
          logs = fetchedLogs ?? [];
          isLoading = false;
        });
      }
    } catch (e) {
      debugPrint("POURQII $e");
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Erreur: $e')),
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: isLoading
          ? const Center(child: CircularProgressIndicator())
          : logs.isEmpty
              ? const Center(
                  child: Text(
                    'Aucune donnée disponible',
                    style: TextStyle(fontSize: 16, color: Colors.grey),
                  ),
                )
              : SingleChildScrollView(
                  child: ConstrainedBox(
                    constraints: BoxConstraints(
                      maxHeight: MediaQuery.of(context).size.height,
                    ),
                    child: CustomDataTable(
                      title: 'Logs des Actions',
                      columns: const [
                        DataColumn(label: Text('Action')),
                        DataColumn(label: Text('Description')),
                        DataColumn(label: Text('Date')),
                      ],
                      rows: logs.map((log) {
                        String formattedDate;

                        if (log['created_at'] != null) {
                          try {
                            DateTime parsedDate =
                                DateTime.parse(log['created_at']);

                            formattedDate = DateFormat('dd/MM/yyyy à HH:mm:ss')
                                .format(parsedDate);
                          } catch (e) {
                            formattedDate = 'Date invalide';
                          }
                        } else {
                          formattedDate = 'Date non disponible';
                        }

                        return DataRow(cells: [
                          DataCell(
                              Text(log['action'] ?? 'Action non spécifiée')),
                          DataCell(Text(log['description'] ??
                              'Description non disponible')),
                          DataCell(Text(formattedDate)),
                        ]);
                      }).toList(),
                      buttonText: '', // Valeur vide
                      onButtonPressed: () {}, // Fonction vide
                    ),
                  ),
                ),
    );
  }
}
