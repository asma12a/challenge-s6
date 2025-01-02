import 'package:flutter/material.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/services/team_service.dart';

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
    if (!_formKey.currentState!.validate()) return;
    try {
      await teamService.updateTeam(widget.eventId, _team);
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
    return Dialog(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Form(
          key: _formKey,
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Text(
                'Modifier l\'équipe',
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                ),
              ),
              TextFormField(
                initialValue: _team.name,
                decoration: InputDecoration(labelText: 'Nom'),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return 'Veuillez entrer un nom';
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
                decoration: InputDecoration(labelText: 'Nombre de joueurs'),
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
                child: Text('Modifier'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
