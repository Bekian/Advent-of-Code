namespace Attempt1;

class Program
{
    static void Main()
    {
        string[] content = [];
        try
        {
            content = File.ReadAllLines("input1.txt");
        }
        catch (Exception ex)
        {
            Console.WriteLine(ex);
        }

        // create a list of lists of strings for the crate values
        // can have a letter if there is a crate, or a space if theres no crate
        List<List<char>> crates = [[]];

        // flag to indicate if we should break out of both loops
        bool shouldBreak = false;
        // this means the first input is over
        for (int i = 0; i < content.Length; i++)
        {
            for (int ii = 1; ii < content[i].Length; ii += 4)
            {
                // stop the loop if the crates section of the input is over
                if (char.IsDigit(content[i][ii]))
                {
                    shouldBreak = true;
                    break;
                }
                // ensure theres enough lists inside the crates list
                if (crates.Count <= content[i].Length % ii)
                {
                    crates.Add([]);
                }

                // add the item to the correct column list 
                crates[content[i].Length % ii].Add(content[i][ii]);
            }
            if (shouldBreak) break;
        }
        foreach (var item in crates)
        {
            foreach (var item2 in item)
            {
                Console.Write(item2);
            }
            Console.WriteLine();
        }
    }
}


// Notes:
/*
~ First task: 
Sort the columns into lists
~ How do i do that?
    get the input @ each space into a list
    ~ How do i do that?
        figure out what each item is
        ~ How do i do that?
            figure out what the possible cases are for a position
            ~ What are those?
                1. a crate
                2. a space
            ~ What do we know about the two possible cases that could help?
                - they're both length of 3, each item is separated by a space
            knowing they're the same length, we can just stream the correct spot as a list item
            ~ How do i do that?
                figure out if theres an easy way to get the values we want
                - we know that each item starts at the second value, this can make the loop simpler
                - we also know that each value is 4 apart from the last so we can increment by 4
                - then our for loop looks like this:
            `for (int i = 1; i < line.Length; i += 4)`
            ~ How do we only get the input we want?
                figure out where the crates input stops
                ~ How do i do that?
                    - we know the crates can be only letters so once we detect a numeric input we stop
Now we have the columns in their initial positions inside of a list of lists
    then map the input lists from rows into columns

*/