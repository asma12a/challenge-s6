import 'package:flutter/material.dart';
import 'package:carousel_slider/carousel_slider.dart';

class Carousel extends StatefulWidget {
  const Carousel({
    super.key,
    required this.imgList,
    required this.text,
  });

  final List<String> imgList;
  final String text;

  @override
  State<Carousel> createState() => CarouselState();
}

class CarouselState extends State<Carousel> {
  int _currentIndex = 0;

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
      builder: (context, constraints) {
        // Calcule la largeur de l'image active (centrée) en fonction de la largeur de l'écran
        double imageWidth = constraints.maxWidth * 0.8;

        return Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Positionne le texte au-dessus de l'image active
            Padding(
              padding: EdgeInsets.only(left: (constraints.maxWidth - imageWidth) / 2),
              child: Text(
                widget.text,
                style: Theme.of(context).textTheme.titleLarge!.copyWith(
                  color: Colors.white
                ),
              ),
            ),
            const SizedBox(height: 10),
            CarouselSlider(
              options: CarouselOptions(
                height: MediaQuery.of(context).size.width,
                autoPlay: true,
                enlargeCenterPage: true,
                onPageChanged: (index, reason) {
                  setState(() {
                    _currentIndex = index;
                  });
                },
              ),
              items: widget.imgList.map((item) {
                return Image.network(
                  item,
                  fit: BoxFit.cover,
                  width: imageWidth,
                );
              }).toList(),
            ),
          ],
        );
      },
    );
  }
}