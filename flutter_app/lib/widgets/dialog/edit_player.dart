import 'package:flutter/material.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/services/team_service.dart';

class EditPlayerDialog extends StatefulWidget {
  final String eventId;
  final Player player;
  final List<Team> teams;
  final Future<void> Function()? onRefresh;

  const EditPlayerDialog({
    super.key,
    required this.eventId,
    required this.player,
    required this.teams,
    this.onRefresh,
  });

  @override
  State<EditPlayerDialog> createState() => _EditPlayerDialogState();
}

class _EditPlayerDialogState extends State<EditPlayerDialog> {
  final TeamService teamService = TeamService();
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();

  late Player _player;

  @override
  void initState() {
    super.initState();

    _player = widget.player;
  }

  void _updatePlayer() async {
    if (_formKey.currentState!.validate()) {
      _formKey.currentState!.save();

      try {
        await teamService.updatePlayer(widget.eventId, _player);
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
              DropdownButtonFormField<String>(
                value: _player.teamID,
                onChanged: (String? value) {
                  setState(() {
                    _player = _player.copyWith(teamID: value);
                  });
                },
                menuMaxHeight: 200,
                items: widget.teams
                    .map((team) => DropdownMenuItem(
                          value: team.id,
                          child: Text(team.name),
                        ))
                    .toList(),
                decoration: const InputDecoration(labelText: 'Équipe'),
              ),
              DropdownButtonFormField<PlayerRole>(
                value: _player.role,
                onChanged: (PlayerRole? value) {
                  if (value == null) return;
                  setState(() {
                    _player = _player.copyWith(role: value);
                  });
                },
                items: PlayerRole.values
                    .map((role) => DropdownMenuItem<PlayerRole>(
                          value: role,
                          child: Text(playerRoleLabel[role] ?? role.name),
                        ))
                    .toList(),
                decoration: const InputDecoration(labelText: 'Rôle'),
              ),
              SizedBox(height: 8),
              ElevatedButton(
                onPressed: _player.role != widget.player.role ||
                        _player.teamID != widget.player.teamID
                    ? _updatePlayer
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
                child: const Text('Modifier'),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
