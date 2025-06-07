# Minecraft Clone

This is a (yet another) Minecraft Clone project, made with golang and raylib for learning purposes. I'm working on this to figure out some basic optimization techniques like meshing and face culling, and also figure out a bit more of 3D computer graphics.

## Resources

- [PlaySpaceFarer's Voxel Meshing Article](https://playspacefarer.com/voxel-meshing/)
- [Perlin Noise](https://rtouti.github.io/graphics/perlin-noise-algorithm)

## TODO

- [ ] extract chunk coords and player coords to a separate struct in a "position" package. Working with vector3 uses float unnecessarily.
- [ ] extract the world loading and unloading into a separate thread.

## Maybe one day

- [ ] replace this perlin noise implementation with a DIY one so I can understand it better.
