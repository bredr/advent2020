import * as fs from "fs";
type Rules = { [k: string]: Array<{ colour: string, count: number }> };
const uniqueBags = (rules: Rules, acc: string[]): string[] => {
  const moreBags = Object.keys(rules).reduce
    ((xx, x) => rules[x].some(b => acc.includes(b.colour)) && !acc.includes(x) ? [...xx, x] : xx, [] as string[]);
  if (moreBags.length === 0) {
    return acc;
  }
  return uniqueBags(rules, [...moreBags, ...acc]);
}


const nestedBags = (rules: Rules, c: string, quantity: number): number => {
  return quantity + rules[c].
    map(({ colour, count }) => nestedBags(rules, colour, count * quantity)).
    reduce((acc, x) => acc + x, 0);
}

try {
  const data = fs.readFileSync('input.txt', 'utf8');
  const input = data.split(/\n/);
  const rules: Rules = input.reduce((acc, x) => {
    const [bagRaw, contains] = x.split("contain");
    const bag = bagRaw.replace(/bags|bag/g, "").trim();
    if (!contains) {
      return acc
    }
    if (contains && contains.includes("no other bags")) {
      return ({ ...acc, [bag]: [] })
    }
    const c = contains.replace(".", "").
      replace(/bags|bag/g, "").
      split(",").
      map(x => x.trim()).
      map(x => {
        const count = x.match(/[(0-9)]+/);
        if (!count) {
          throw Error();
        }
        const colour = x.replace(count[0], "").trim()
        return { colour, count: parseInt(count[0]) };
      });
    return ({ ...acc, [bag]: c })
  }, {})
  const result1 = uniqueBags(rules, ["shiny gold"]);
  console.log("Result part1 =", result1.length - 1);

  const result2 = nestedBags(rules, "shiny gold", 1);
  console.log("Result part2 =", result2 - 1);

} catch (e) {
  console.error(e);
}
