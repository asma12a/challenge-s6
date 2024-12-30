import 'package:flutter/material.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/services/team_service.dart';

class EditTeamDialog extends StatefulWidget {
  final Team team;
  final Future<void> Function()? onRefresh;

  const EditTeamDialog({super.key, required this.team, this.onRefresh});

  @override
  State<EditTeamDialog> createState() => _EditTeamDialogState();
}

class _EditTeamDialogState extends State<EditTeamDialog> {
  final TeamService eventService = TeamService();

  @override
  Widget build(BuildContext context) {
    return const Placeholder();
  }
}
