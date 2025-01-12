import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class OfflineDialog extends StatelessWidget {
  const OfflineDialog({super.key});

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);

    return Dialog(
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            const Icon(
              Icons.signal_wifi_off,
              size: 48,
              color: Colors.grey,
            ),
            Text(
              translate?.no_internet ?? 'Aucune connexion internet',
              style: const TextStyle(
                fontSize: 18,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: 16),
            Text(
              translate?.check_internet ??
                  'Veuillez vérifier votre connexion internet et réessayer.',
              style: const TextStyle(
                fontSize: 16,
              ),
            ),
            const SizedBox(height: 16),
            ElevatedButton(
              onPressed: () {
                Navigator.of(context).pop();
              },
              child: const Text('OK'),
            ),
          ],
        ),
      ),
    );
  }
}
