import random
import string
import time


class Robot:
    def __init__(self):
        self.letters = list(string.ascii_uppercase)
        self.digits = list(string.digits)
        self.name = self.robot_name()

    def robot_name(self):
        self.random()
        return f"{''.join(random.sample(k=2,population=self.letters))}{''.join(random.sample(k=3,population=self.digits))}"

    def reset(self):
        self.random()
        self.name = self.robot_name()
        return self.name

    def random(self):
        random.seed(time.time())


if __name__ == '__main__':
    # Call the generator
    seed = 1234
    random.seed(seed)
    robot = Robot()
    name = robot.name

    # Reinitialize RNG using seed
    random.seed(seed)

    # Call the generator again
    robot.reset()
    name2 = robot.name
    print(name)
    print(name2)
