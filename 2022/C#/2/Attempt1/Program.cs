using System.Runtime.CompilerServices;

namespace Attempt1;

class Program
{
    static void Main(string[] args)
    {
        string[] content = [];
        try
        {
            content = File.ReadAllLines("input.txt");
        }
        catch (Exception ex)
        {
            Console.WriteLine(ex);
        }

        // define the map to hold all possible combinations
        string[] combinations = ["A X", "A Y", "A Z", "B X", "B Y", "B Z", "C X", "C Y", "C Z"];
        int[] scores = [4, 8, 3, 1, 5, 9, 7, 2, 6];
        int player1Score = 0;
        // PART 2:
        // the same but the second column means something different:
        // lose X, draw Y, win Z
        // so we need to just make a new array of scores:
        // basically use the columns to determine which symbol you use
        // ay is 4, bx is 1, cz is 7
        int[] part2Scores = [3, 4, 8, 1, 5, 9, 2, 6, 7];
        int part2Score = 0;

        // check each line in the input
        foreach (var line in content)
        {
            // and check each line with each combination
            foreach (var combination in combinations)
            {
                // if the line is equal to the current combination
                if (line == combination)
                {
                    // get the index of the current combination in combinations
                    var indexOfCombination = Array.IndexOf(combinations, combination);
                    // then use that index to reference the corresponding item in the scores and add that to the scores
                    player1Score += scores[indexOfCombination];
                    part2Score += part2Scores[indexOfCombination];
                    // then break the first loop since we know the combinations wont match the remaing ones
                    break;
                }
                // otherwise continue;
            }
        }
        Console.WriteLine($"total score {player1Score}, Part 2 Score: {part2Score}");
    }
}
