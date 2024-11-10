import 'package:flutter/material.dart';
import 'package:flutter_app/screens/home.dart';
import 'package:flutter_app/screens/search.dart';
import 'package:flutter_app/widgets/main_drawer.dart';

class TabsScreen extends StatefulWidget {
  const TabsScreen({super.key});

  @override
  State<TabsScreen> createState() => _TabsScreenState();
}

class _TabsScreenState extends State<TabsScreen> {
  int _selectPageIndex = 0;

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
    Widget activePage = HomeScreen();

    if (_selectPageIndex == 1) {
      activePage = SearchScreen();
    }

    return Scaffold(
      appBar: AppBar(
        actions: [
          if (activePage.runtimeType == HomeScreen)
            IconButton(
              icon: const Icon(Icons.add),
              onPressed: () {
                // Action pour l'icône "add"
              },
            ),
        ],
      ),
      drawer: MainDrawer(onSelectScreen: _setScreen),
      body: Padding(padding: EdgeInsets.only(top: 20), child: activePage),
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
