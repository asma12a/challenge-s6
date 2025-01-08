import 'package:flutter/material.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/services/team_service.dart';
import 'package:squad_go/main.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class AddPlayerDialog extends StatefulWidget {
  final String eventId;
  final String teamId;
  final Future<void> Function()? onRefresh;

  const AddPlayerDialog({
    super.key,
    required this.eventId,
    required this.teamId,
    this.onRefresh,
  });

  @override
  State<AddPlayerDialog> createState() => _AddPlayerDialogState();
}

class _AddPlayerDialogState extends State<AddPlayerDialog> {
  final TeamService teamService = TeamService();
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();

  PlayerRole _selectedRole = PlayerRole.player;
  String? _email;

  void _addPlayer() async {
    if (_email == null) return;

    try {
      await teamService.addPlayerToTeam(
          widget.eventId, widget.teamId, _email!, _selectedRole);
      widget.onRefresh?.call();
      Navigator.of(context).pop();
    } on AppException catch (e) {
      // Handle AppException error
      log.severe('Failed to add player to team: ${e.message}');
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Erreur: ${e.message}')),
      );
    } catch (e) {
      // Handle other errors
      log.severe('Failed to add player to team: $e');
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Une erreur est survenue')),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);
    final translate = AppLocalizations.of(context);
    return Dialog(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Form(
          key: _formKey,
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Text(translate?.add_player ?? 
              Text(translate?.add_player ?? 
                'Ajouter un joueur',
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
                style: TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                ),
              ),
              SizedBox(height: 16),
              TextFormField(
                decoration: const InputDecoration(labelText: 'Email'),
                keyboardType: TextInputType.emailAddress,
                validator: (value) {
                  // add email validation
                  if (value == null || value.isEmpty) {
                    return 'Veuillez entrer une adresse email';
                  }

                  bool emailValid = RegExp(
                          r"^[a-zA-Z0-9.a-zA-Z0-9.!#$%&'*+-/=?^_`{|}~]+@[a-zA-Z0-9]+\.[a-zA-Z]+")
                      .hasMatch(value);
                  if (!emailValid) {
                    return 'Veuillez entrer une adresse email valide';
                  }

                  return null;
                },
                onChanged: (value) {
                  setState(() {
                    _email = value;
                  });
                },
                onTapOutside: (event) {
                  FocusScope.of(context).unfocus();
                },
              ),
              SizedBox(height: 16),
              DropdownButtonFormField<PlayerRole>(
                value: _selectedRole,
                onChanged: (PlayerRole? value) {
                  if (value == null) return;
                  setState(() {
                    _selectedRole = value;
                  });
                },
                items: PlayerRole.values
                    .map((role) => DropdownMenuItem<PlayerRole>(
                          value: role,
                          child: Text(playerRoleLabel[role] ?? role.name),
                        ))
                    .toList(),
                decoration: const InputDecoration(labelText: 'Role'),
              ),
              SizedBox(height: 16),
              ElevatedButton.icon(
                onPressed: _formKey.currentState?.validate() ?? false
                    ? _addPlayer
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
                icon: const Icon(
                  Icons.add,
                  color: Colors.white,
                ),
                label: Text(translate?.add ?? 'Ajouter'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
