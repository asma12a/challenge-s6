import 'package:flutter/material.dart';
import 'package:squad_go/core/services/log_service.dart';
import 'package:squad_go/core/models/action_log.dart';
import 'package:squad_go/platform/web/screens/custom_data_table.dart';
import 'package:intl/intl.dart';  // Importez la bibliothèque intl

class AdminLogsPage extends StatefulWidget {
  const AdminLogsPage({super.key});

  @override
  State<AdminLogsPage> createState() => _AdminLogsPageState();
}

class _AdminLogsPageState extends State<AdminLogsPage> {
  late Future<List<ActionLog>> _logsFuture;

  @override
  void initState() {
    super.initState();
    _loadLogs();
  }

  void _loadLogs() {
    setState(() {
      _logsFuture = LogService.getLogs();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: FutureBuilder<List<ActionLog>>(
        future: _logsFuture,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          } else if (snapshot.hasError) {
            return Center(
              child: Text(
                'Erreur: ${snapshot.error}',
                style: TextStyle(color: Theme.of(context).colorScheme.error),
              ),
            );
          } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
            return const Center(child: Text('Aucun log trouvé.'));
          }

          final logs = snapshot.data!;

          return SingleChildScrollView(
            child: CustomDataTable(
              title: 'Logs des Actions',
              columns: [
                DataColumn(label: Text('ID')),
                DataColumn(label: Text('Utilisateur')),
                DataColumn(label: Text('Action')),
                DataColumn(label: Text('Description')),
                DataColumn(label: Text('Date')),
              ],
              rows: logs.map((log) {
                String formattedDate = DateFormat('dd/MM/yyyy à HH:mm:ss').format(log.createdAt);

                return DataRow(cells: [
                  DataCell(Text(log.id)),
                  DataCell(Text(log.userId)),
                  DataCell(Text(log.action)),
                  DataCell(Text(log.description)),
                  DataCell(Text(formattedDate)), 
                ]);
              }).toList(),
              buttonText: 'Ajouter un Log',
              onButtonPressed: () {
              },
            ),
          );
        },
      ),
    );
  }
}
