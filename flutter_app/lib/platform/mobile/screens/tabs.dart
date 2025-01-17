import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/providers/connectivity_provider.dart';
import 'package:squad_go/platform/mobile/screens/account.dart';
import 'package:squad_go/platform/mobile/screens/join.dart';
import 'package:squad_go/platform/mobile/screens/home.dart';
import 'package:squad_go/platform/mobile/screens/search.dart';
import 'package:squad_go/platform/mobile/widgets/dialog/offline.dart';
import 'package:squad_go/platform/mobile/widgets/main_drawer.dart';
import 'package:squad_go/platform/mobile/screens/new_event.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class TabsScreen extends StatefulWidget {
  final int? initialPageIndex;
  final bool? shouldRefresh;

  const TabsScreen({super.key, this.initialPageIndex, this.shouldRefresh});

  @override
  State<TabsScreen> createState() => _TabsScreenState();
}

class _TabsScreenState extends State<TabsScreen> {
  int _selectPageIndex = 0;

  @override
  void initState() {
    super.initState();
    if (widget.initialPageIndex != null) {
      _selectPageIndex = widget.initialPageIndex!;
    }
  }

  void _selectPage(int index) {
    setState(() {
      _selectPageIndex = index;
    });
  }

  void _setScreen(String identifier) async {
    Navigator.of(context).pop();
  }

  @override
  Widget build(BuildContext context) {
    final translate = AppLocalizations.of(context);
    Widget activePage = HomeScreen();

    if (_selectPageIndex == 1) {
      activePage = SearchScreen();
    }

    if (_selectPageIndex == 2) {
      activePage = JoinEventScreen();
    }

    if (_selectPageIndex == 3) {
      activePage = AccountScreen();
    }

    var isOnline = context.watch<ConnectivityState>().isConnected;
    WidgetsBinding.instance.addPostFrameCallback((_) {
      ScaffoldMessenger.of(context).clearSnackBars();
    });

    if (!isOnline) {
      WidgetsBinding.instance.addPostFrameCallback((_) {
        ScaffoldMessenger.of(context).clearSnackBars();

        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Icon(Icons.wifi_off,
                    color: Theme.of(context).colorScheme.onPrimary),
                const SizedBox(width: 10),
                Text(
                  translate?.no_internet ??
                      "Vous n'êtes pas connecté à internet.",
                  style:
                      TextStyle(color: Theme.of(context).colorScheme.onPrimary),
                ),
              ],
            ),
            backgroundColor: Colors.grey.shade700,
            duration: Duration(days: 5),
            dismissDirection: DismissDirection.none,
          ),
        );
      });
    }

    return Scaffold(
      appBar: AppBar(
        actions: [
          if (activePage.runtimeType == HomeScreen)
            IconButton(
              icon: const Icon(Icons.add),
              onPressed: () {
                if (!isOnline) {
                  showDialog(
                    context: context,
                    builder: (context) => const OfflineDialog(),
                  );
                  return;
                }

                Navigator.of(context).push(
                  MaterialPageRoute(
                    builder: (ctx) => NewEvent(),
                  ),
                );
              },
            ),
        ],
      ),
      drawer: MainDrawer(onSelectScreen: _setScreen),
      body: activePage,
      bottomNavigationBar: BottomNavigationBar(
        type: BottomNavigationBarType.fixed,
        onTap: _selectPage,
        currentIndex: _selectPageIndex,
        items: [
          BottomNavigationBarItem(
            icon: const Icon(Icons.home),
            label: translate?.tabs_home ?? 'Accueil',
          ),
          BottomNavigationBarItem(
            icon: const Icon(Icons.search),
            label: translate?.tabs_search ?? 'Rechercher',
          ),
          BottomNavigationBarItem(
            icon: const Icon(Icons.qr_code),
            label: translate?.tabs_join ?? 'Rejoindre',
          ),
          BottomNavigationBarItem(
            icon: const Icon(Icons.person),
            label: translate?.tabs_profile ?? 'Profile',
          ),
        ],
      ),
    );
  }
}
