import 'package:flutter/material.dart';

class EditStatsPlayerDialog extends StatefulWidget {
  final Future<void> Function()? onRefresh;

  const EditStatsPlayerDialog({super.key, this.onRefresh});

  @override
  State<EditStatsPlayerDialog> createState() => _EditStatsPlayerDialogState();
}

class _EditStatsPlayerDialogState extends State<EditStatsPlayerDialog> {
  @override
  Widget build(BuildContext context) {
    return const Placeholder();
  }
}
