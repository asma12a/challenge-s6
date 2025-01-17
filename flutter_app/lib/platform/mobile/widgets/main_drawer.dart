import 'package:flutter/material.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/core/providers/locale_provider.dart';
import 'package:squad_go/platform/mobile/widgets/logo.dart';
import 'package:provider/provider.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class MainDrawer extends StatelessWidget {
  const MainDrawer({super.key, required this.onSelectScreen});

  final void Function(String identifier) onSelectScreen;
  final storage = const FlutterSecureStorage();

  void _logOut(BuildContext context) async {
    await context.read<AuthState>().logout();
    context.go('/sign-in');
  }

  void _changeLanguage(BuildContext context, String languageCode) {
    final localeProvider = context.read<LocaleProvider>();
    localeProvider.setLocale(Locale(languageCode));
  }

  @override
  Widget build(BuildContext context) {
    final userInfo = context.read<AuthState>().userInfo;
    final translate = AppLocalizations.of(context);
    return Drawer(
      child: Column(
        children: [
          DrawerHeader(
            padding: const EdgeInsets.all(20),
            decoration: BoxDecoration(
              gradient: LinearGradient(
                colors: [
                  Theme.of(context).colorScheme.primaryContainer,
                  Theme.of(context)
                      .colorScheme
                      .primaryContainer
                      .withOpacity(0.8),
                ],
                begin: Alignment.topLeft,
                end: Alignment.bottomRight,
              ),
            ),
            child: Row(
              children: [
                Logo(width: 120),
                const SizedBox(width: 18),
              ],
            ),
          ),
          ListTile(
            leading: Icon(
              Icons.person,
              size: 26,
              color: Theme.of(context).colorScheme.onSurface,
            ),
            title: Text(
              userInfo!.name,
              style: Theme.of(context).textTheme.titleSmall!.copyWith(
                    color: Theme.of(context).colorScheme.onSurface,
                    fontSize: 24,
                  ),
            ),
            onTap: () {
              onSelectScreen('meals');
            },
          ),
          SizedBox(height: 20),
          // Ajout du Dropdown pour changer la langue
          ListTile(
            leading: Icon(
              Icons.language,
              size: 26,
              color: Theme.of(context).colorScheme.onSurface,
            ),
            title: DropdownButton<String>(
              value: Localizations.localeOf(context)
                  .languageCode, // Valeur actuelle de la langue
              onChanged: (String? newLanguage) {
                if (newLanguage != null) {
                  _changeLanguage(context, newLanguage);
                }
              },
              items: [
                DropdownMenuItem(
                  value: 'en',
                  child: Text(translate?.english ??
                    'English',
                    style: TextStyle(fontSize: 24),
                  ),
                ),
                DropdownMenuItem(
                  value: 'fr',
                  child: Text(translate?.french ??
                    'Français',
                    style: TextStyle(fontSize: 24),
                  ),
                ),
                DropdownMenuItem(
                  value: 'es',
                  child: Text(translate?.spanish ??
                    'Espagnol',
                    style: TextStyle(fontSize: 24),
                  ),
                ),
                DropdownMenuItem(
                  value: 'de',
                  child: Text(translate?.german ??
                    'Allemand',
                    style: TextStyle(fontSize: 24),
                  ),
                ),
              ],
            ),
          ),
          SizedBox(height: 250),
          Expanded(
            child: Align(
              child: ListTile(
                leading: Icon(
                  Icons.logout,
                  size: 26,
                  color: Theme.of(context).colorScheme.onSurface,
                ),
                title: Text(
                  translate?.logout ?? 'Déconnexion',
                  style: Theme.of(context).textTheme.titleSmall!.copyWith(
                        color: Theme.of(context).colorScheme.onSurface,
                        fontSize: 24,
                      ),
                ),
                onTap: () => _logOut(context),
              ),
            ),
          )
        ],
      ),
    );
  }
}
