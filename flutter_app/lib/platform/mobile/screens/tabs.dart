import 'package:flutter/material.dart';
import 'package:squad_go/platform/mobile/screens/account.dart';
import 'package:squad_go/platform/mobile/screens/join.dart';
import 'package:squad_go/platform/mobile/screens/home.dart';
import 'package:squad_go/platform/mobile/screens/search.dart';
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

    return Scaffold(
      appBar: AppBar(
        actions: [
          if (activePage.runtimeType == HomeScreen)
            IconButton(
              icon: const Icon(Icons.add),
              onPressed: () {
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
            icon: Icon(Icons.home),
            label: translate?.tabs.home ?? 'Accueil',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.search),
            label: translate?.tabs.search ?? 'Rechercher',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.qr_code),
            label: translate?.tabs.join ?? 'Rejoindre',
          ),
          const BottomNavigationBarItem(
            icon: Icon(Icons.person),
            label: 'Profile',
          ),
        ],
      ),
    );
  }
}
