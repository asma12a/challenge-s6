import 'package:flutter/material.dart';

// TODO: Full details of player: player info + player stats
class ShowPlayerDetailsDialog extends StatefulWidget {
  final Future<void> Function()? onRefresh;

  const ShowPlayerDetailsDialog({super.key, this.onRefresh});

  @override
  State<ShowPlayerDetailsDialog> createState() =>
      _ShowPlayerDetailsDialogState();
}

class _ShowPlayerDetailsDialogState extends State<ShowPlayerDetailsDialog> {
  @override
  Widget build(BuildContext context) {
    return const Placeholder();
  }
}
