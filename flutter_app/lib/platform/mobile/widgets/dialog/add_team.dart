import 'package:flutter/material.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/services/team_service.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class AddTeamDialog extends StatefulWidget {
  final String eventId;
  final Future<void> Function()? onRefresh;

  const AddTeamDialog({super.key, required this.eventId, this.onRefresh});

  @override
  State<AddTeamDialog> createState() => _AddTeamDialogState();
}

class _AddTeamDialogState extends State<AddTeamDialog> {
  final TeamService teamService = TeamService();
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();

  String? _teamName;
  int? _maxPlayers;

  void _createTeam() async {
    if (_teamName == null || !_formKey.currentState!.validate()) return;
    try {
      await teamService.createTeam(widget.eventId, _teamName!, _maxPlayers);
      widget.onRefresh?.call();
      Navigator.of(context).pop();
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
  }

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);
    return Dialog(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Form(
          key: _formKey,
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Text(translate?.new_team ??
                'Nouvelle équipe',
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                ),
              ),
              SizedBox(height: 16),
              TextFormField(
                decoration: InputDecoration(labelText: 'Nom de l\'équipe'),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Veuillez entrer un nom d\'équipe';
                  }
                  return null;
                },
                onChanged: (value) => _teamName = value,
                onTapOutside: (event) {
                  FocusScope.of(context).unfocus();
                },
              ),
              TextFormField(
                decoration: InputDecoration(labelText: translate?.nb_players ?? 'Nombre de joueurs'),
                keyboardType: TextInputType.number,
                onChanged: (value) => _maxPlayers = int.tryParse(value),
                onTapOutside: (event) {
                  FocusScope.of(context).unfocus();
                },
              ),
              SizedBox(height: 16),
              ElevatedButton(
                onPressed: _formKey.currentState?.validate() ?? false
                    ? _createTeam
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
                child: Text('Créer l\'équipe'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
