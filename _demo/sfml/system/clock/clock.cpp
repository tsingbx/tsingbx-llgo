#include <SFML/System/Clock.hpp>

int main(void) {
    sf::Clock clock;
    clock.getElapsedTime().asSeconds();
    return 0;
}