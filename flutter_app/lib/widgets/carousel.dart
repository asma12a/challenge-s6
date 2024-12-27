import 'package:flutter/material.dart';
import 'package:carousel_slider/carousel_slider.dart';

class Carousel extends StatefulWidget {
  const Carousel({
    super.key,
    required this.items,
    required this.text,
    this.isLoading = false,
  });

  final List<Widget> items;
  final String text;
  final bool isLoading;

  @override
  State<Carousel> createState() => CarouselState();
}

class CarouselState extends State<Carousel> {
  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
      builder: (context, constraints) {
        return Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Padding(
              padding: const EdgeInsets.only(left: 40),
              child: Text(
                widget.text,
                style: Theme.of(context).textTheme.titleLarge!,
              ),
            ),
            const SizedBox(
              height: 10,
            ),
            if (widget.isLoading)
              Padding(
                padding: const EdgeInsets.only(left: 40),
                child: const Center(
                  child: CircularProgressIndicator(),
                ),
              ),
            if (!widget.isLoading && widget.items.isEmpty)
              Padding(
                padding: const EdgeInsets.only(left: 40),
                child: Text('Aucun élément à afficher'),
              ),
            CarouselSlider(
              options: CarouselOptions(
                enableInfiniteScroll: false,
                padEnds: false,
                height: widget.items.isEmpty ? 0 : 20,
              ),
              items: widget.items,
            ),
            const SizedBox(
              height: 20,
            ),
          ],
        );
      },
    );
  }
}
