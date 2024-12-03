namespace Attempt1;
// this attmept converts the input to an array
// then logically compares the bounds in the row to determine if they overlap

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

        // create a list of lists of ints for the positions
        List<List<int>> positions = [[]];

        // assign the positions
        for (int i = 0; i < content.Length; i++)
        {
            string[] parts = content[i].Split(',');
            string[] firstParts = parts[0].Split('-');
            string[] secondParts = parts[1].Split('-');

            positions.Add([int.Parse(firstParts[0]), int.Parse(firstParts[1]), int.Parse(secondParts[0]), int.Parse(secondParts[1])]);
        }
        positions.RemoveAt(0);
        int overlapCount = 0;
        int overlapCount2 = 0;
        // now we check if each pair of elves sections completely overlap
        foreach (var row in positions)
        {
            // first check if complete overlap
            // check if first elf surrounds second
            if (row[0] <= row[2] && row[1] >= row[3])
            {
                overlapCount++;
            }
            // check if second elf surrounds first
            else if (row[2] <= row[0] && row[3] >= row[1])
            {
                overlapCount++;
            }

            // // Part 2
            // check if first start is less than or equal to the second end
            // AND if second start is less than or equal to the first end
            // then we know the bounds overlap
            if (row[0] <= row[3] && row[2] <= row[1])
            {
                overlapCount2++;
            }
        }
        Console.WriteLine($"{overlapCount}, {overlapCount2}");
    }
}

// 577 too high
// 550 is correct for part 1

// first try at part 2 is 960, too high
// second try at part 2   791, too low
// third try 931 is correct