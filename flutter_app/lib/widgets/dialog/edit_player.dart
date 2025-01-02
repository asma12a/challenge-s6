import 'package:flutter/material.dart';

// TODO: Can edit player team + role
class EditPlayerDialog extends StatefulWidget {
  final Future<void> Function()? onRefresh;

  const EditPlayerDialog({super.key, this.onRefresh});

  @override
  State<EditPlayerDialog> createState() => _EditPlayerDialogState();
}

class _EditPlayerDialogState extends State<EditPlayerDialog> {
  @override
  Widget build(BuildContext context) {
    return const Placeholder();
  }
}
