import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:squad_go/core/models/event.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';


class ShareEventDialog extends StatefulWidget {
  final Event event;
  const ShareEventDialog({super.key, required this.event});

  @override
  State<ShareEventDialog> createState() => _ShareEventDialogState();
}

class _ShareEventDialogState extends State<ShareEventDialog> {
  final String appUrl =
      String.fromEnvironment('APP_URL', defaultValue: 'https://squad-go.com');

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);

    return Dialog(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Text(translate?.share_event ??
              'Partager l\'évent',
              style: Theme.of(context).textTheme.headlineSmall,
            ),
            const SizedBox(height: 16),
            TextField(
              readOnly: true,
              decoration: InputDecoration(
                labelText: translate?.event_code ?? 'Code de l\'évent',
                border: OutlineInputBorder(),
                suffixIcon: IconButton(
                  icon: const Icon(Icons.copy),
                  onPressed: () {
                    Clipboard.setData(ClipboardData(text: widget.event.code));
                    ScaffoldMessenger.of(context).showSnackBar(
                      SnackBar(content: Text(translate?.copy ?? 'Copié dans le presse-papier')),
                    );
                  },
                ),
              ),
              controller: TextEditingController(text: widget.event.code),
            ),
            const SizedBox(height: 16),
            TextField(
              readOnly: true,
              decoration: InputDecoration(
                labelText: translate?.event_link ?? 'Lien de l\'évent',
                border: OutlineInputBorder(),
                suffixIcon: IconButton(
                  icon: const Icon(Icons.copy),
                  onPressed: () {
                    Clipboard.setData(ClipboardData(
                        text: '$appUrl/events/${widget.event.id!}'));
                    ScaffoldMessenger.of(context).showSnackBar(
                      SnackBar(content: Text(translate?.copy ?? 'Copié dans le presse-papier')),
                    );
                  },
                ),
              ),
              controller: TextEditingController(
                  text: '$appUrl/events/${widget.event.id!}'),
            ),
            const SizedBox(height: 16),
            ElevatedButton(
              onPressed: () => Navigator.pop(context),
              child: Text(translate?.cancel ?? 'Fermer'),
            ),
          ],
        ),
      ),
    );
  }
}
