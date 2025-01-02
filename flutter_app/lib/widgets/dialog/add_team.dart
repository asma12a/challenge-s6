import 'package:flutter/material.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/services/team_service.dart';

class AddTeamDialog extends StatefulWidget {
  final Future<void> Function()? onRefresh;

  const AddTeamDialog({super.key, this.onRefresh});

  @override
  State<AddTeamDialog> createState() => _AddTeamDialogState();
}

class _AddTeamDialogState extends State<AddTeamDialog> {
  final TeamService eventService = TeamService();

  Team? team;

  @override
  Widget build(BuildContext context) {
    return const Placeholder();
  }
}
