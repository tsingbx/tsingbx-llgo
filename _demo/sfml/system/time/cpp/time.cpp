#include <SFML/System/Time.hpp>
#include <stdio.h>
int main(void) {
    sf::Time time = sf::Time();
    printf("time.asSeconds:%f\n", time.asSeconds());
    printf("time.asMilliseconds:%d\n", time.asMilliseconds());
    printf("time.asMicroseconds:%ld\n", time.asMicroseconds());
    printf("sf::Time::Zero.asSeconds:%f\n", sf::Time::Zero.asSeconds());
    printf("sf::Time::Zero.asMilliseconds:%d\n", sf::Time::Zero.asMilliseconds());
    printf("sf::Time::Zero.asMicroseconds:%ld\n", sf::Time::Zero.asMicroseconds());
    return 0;
}