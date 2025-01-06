import 'package:flutter/material.dart';
import 'package:squad_go/platform/mobile/screens/join.dart';
import 'package:squad_go/platform/mobile/screens/home.dart';
import 'package:squad_go/platform/mobile/screens/search.dart';
import 'package:squad_go/platform/mobile/widgets/main_drawer.dart';
import 'package:squad_go/platform/mobile/screens/new_event.dart';

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
    Widget activePage = HomeScreen(shouldRefresh: widget.shouldRefresh);

    if (_selectPageIndex == 1) {
      activePage = SearchScreen();
    }

    if (_selectPageIndex == 2) {
      activePage = JoinEventScreen();
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
        items: const [
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: 'Accueil',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.search),
            label: 'Rechercher',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.qr_code),
            label: 'Rejoindre',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.event_available),
            label: 'Mes évents',
          ),
        ],
      ),
    );
  }
}
