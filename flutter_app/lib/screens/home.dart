import 'package:flutter/material.dart';
import 'package:squad_go/widgets/home_widgets/my_events.dart';
import 'package:squad_go/widgets/home_widgets/recommended_events.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => HomeScreenState();
}

class HomeScreenState extends State<HomeScreen>
    with SingleTickerProviderStateMixin {
  late AnimationController _animationController;
  final GlobalKey<HomeRecommendedEventsState> recommendedEventsKey =
      GlobalKey<HomeRecommendedEventsState>();
  final GlobalKey<HomeMyEventsState> myEventsKey =
      GlobalKey<HomeMyEventsState>();

  @override
  void initState() {
    super.initState();

    _animationController = AnimationController(
      vsync: this,
      duration: const Duration(milliseconds: 300),
      lowerBound: 0,
      upperBound: 1,
    );

    _animationController.forward();
  }

  @override
  void dispose() {
    _animationController.dispose();
    super.dispose();
  }

  Future<void> _refresh() async {
    await myEventsKey.currentState?.fetchMyEvents();
    await recommendedEventsKey.currentState?.fetchRecommendedEvents();
  }

  @override
  Widget build(BuildContext context) {
    return AnimatedBuilder(
      builder: (context, child) => SlideTransition(
        position: Tween(
          begin: const Offset(0, 0.3),
          end: const Offset(0, 0),
        ).animate(
          CurvedAnimation(
            parent: _animationController,
            curve: Curves.easeInOut,
          ),
        ),
        child: child,
      ),
      animation: _animationController,
      child: RefreshIndicator(
        onRefresh: _refresh,
        child: SingleChildScrollView(
          physics: const AlwaysScrollableScrollPhysics(),
          child: Column(
            children: [
              const SizedBox(
                height: 10,
              ),
              HomeMyEvents(key: myEventsKey),
              HomeRecommendedEvents(key: recommendedEventsKey)
            ],
          ),
        ),
      ),
    );
  }
}
