import java.util.Scanner;

public class BitcoinFarmGame {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int bitcoin = 0;
        int farms = 0;
        int animals = 0;

        System.out.println("Welcome to Bitcoin Farm Game!");
        while (true) {
            System.out.println("\nMenu:");
            System.out.println("1. Plant Bitcoin Seeds (+10 Bitcoin)");
            System.out.println("2. Buy a Farm (50 Bitcoin)");
            System.out.println("3. Buy Animals (100 Bitcoin)");
            System.out.println("4. Exit");
            System.out.print("Choose an action: ");
            int choice = scanner.nextInt();

            switch (choice) {
                case 1:
                    bitcoin += 10;
                    System.out.println("You planted Bitcoin seeds! Current Bitcoin: " + bitcoin);
                    break;
                case 2:
                    if (bitcoin >= 50) {
                        farms++;
                        bitcoin -= 50;
                        System.out.println("You bought a farm! Total Farms: " + farms);
                    } else {
                        System.out.println("Not enough Bitcoin to buy a farm!");
                    }
                    break;
                case 3:
                    if (bitcoin >= 100) {
                        animals++;
                        bitcoin -= 100;
                        System.out.println("You bought animals! Total Animals: " + animals);
                    } else {
                        System.out.println("Not enough Bitcoin to buy animals!");
                    }
                    break;
                case 4:
                    System.out.println("Thanks for playing! Final Stats: Farms=" + farms + ", Animals=" + animals);
                    scanner.close();
                    return;
                default:
                    System.out.println("Invalid choice. Try again.");
            }
        }
    }
}
