import 'package:flutter/material.dart';

class AddPlayerDialog extends StatefulWidget {
  final Future<void> Function()? onRefresh;

  const AddPlayerDialog({super.key, this.onRefresh});

  @override
  State<AddPlayerDialog> createState() => _AddPlayerDialogState();
}

class _AddPlayerDialogState extends State<AddPlayerDialog> {
  @override
  Widget build(BuildContext context) {
    return const Placeholder();
  }
}
