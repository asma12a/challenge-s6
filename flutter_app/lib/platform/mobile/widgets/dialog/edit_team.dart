import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/providers/connectivity_provider.dart';
import 'package:squad_go/core/services/team_service.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/offline.dart';

class EditTeamDialog extends StatefulWidget {
  final String eventId;
  final Team team;
  final Future<void> Function()? onRefresh;

  const EditTeamDialog({
    super.key,
    required this.team,
    required this.eventId,
    this.onRefresh,
  });

  @override
  State<EditTeamDialog> createState() => _EditTeamDialogState();
}

class _EditTeamDialogState extends State<EditTeamDialog> {
  final TeamService teamService = TeamService();
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();
  late Team _team;

  @override
  void initState() {
    super.initState();
    _team = widget.team;
  }

  void _updateTeam() async {
    final translate = AppLocalizations.of(context);

    if (!_formKey.currentState!.validate()) return;
    try {
      await teamService.updateTeam(widget.eventId, _team);
      widget.onRefresh?.call();
      Navigator.of(context).pop();
    } on AppException catch (e) {
      // Handle AppException error
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text(
            '${translate?.error ?? "Erreur:"} ${e.message}',
          ),
        ),
      );
    } catch (e) {
      // Handle other errors
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text(translate?.error_occurred ?? 'Une erreur est survenue')),
      );
    }
  }

  void _deleteTeam() async {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: Text('Supprimer l\'équipe'),
          content: Text(
              'Êtes-vous sûr de vouloir supprimer l\'équipe ${_team.name} ?'),
          actions: [
            TextButton(
              child: Text("Annuler"),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
            TextButton(
              child: Text("OK"),
              onPressed: () async {
                try {
                  await teamService.deleteTeam(widget.eventId, _team.id);
                  widget.onRefresh?.call();
                } on AppException catch (e) {
                  // Handle AppException error
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(content: Text('Erreur: ${e.message}')),
                  );
                } catch (e) {
                  // Handle other errors
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(content: Text('Une erreur est survenue')),
                  );
                }
                Navigator.of(context).pop();
                Navigator.of(context).pop();
              },
            ),
          ],
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);

    var isOnline = context.watch<ConnectivityState>().isConnected;

    return Dialog(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Form(
          key: _formKey,
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Text(
                '${translate?.edit ?? "Modifier:"} ${_team.name}',
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                ),
              ),
              TextFormField(
                initialValue: _team.name,
                decoration: InputDecoration(labelText: translate?.name ?? 'Nom'),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return translate?.empty_team_name ?? 'Veuillez entrer un nom';
                  }
                  return null;
                },
                onChanged: (value) {
                  _team = _team.copyWith(name: value);
                },
                onTapOutside: (event) {
                  FocusScope.of(context).unfocus();
                },
              ),
              TextFormField(
                initialValue: _team.maxPlayers.toString(),
                decoration: InputDecoration(
                    labelText: translate?.nb_players ?? 'Nombre de joueurs'),
                keyboardType: TextInputType.number,
                onChanged: (value) {
                  _team = _team.copyWith(maxPlayers: int.tryParse(value) ?? 0);
                },
                onTapOutside: (event) {
                  FocusScope.of(context).unfocus();
                },
              ),
              SizedBox(height: 16),
              ElevatedButton(
                onPressed: _team.name != widget.team.name ||
                        _team.maxPlayers != widget.team.maxPlayers
                    ? _updateTeam
                    : null,
                style: ElevatedButton.styleFrom(
                  foregroundColor: Colors.white,
                  backgroundColor: Colors.blue,
                  padding:
                      const EdgeInsets.symmetric(horizontal: 20, vertical: 10),
                  textStyle: const TextStyle(
                    fontSize: 16,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                child: Text(translate?.edit ?? 'Modifier'),
              ),
              TextButton(
                onPressed: () {
                  if (!isOnline) {
                    showDialog(
                      context: context,
                      builder: (context) => const OfflineDialog(),
                    );
                    return;
                  }
                  _deleteTeam();
                },
                style: TextButton.styleFrom(
                  foregroundColor: Colors.red.shade300,
                  textStyle: const TextStyle(
                    fontSize: 16,
                    fontWeight: FontWeight.bold,
                    decoration: TextDecoration.underline,
                  ),
                ),
                child: const Text('Supprimer l\'équipe'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
