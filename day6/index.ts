import * as fs from "fs";



try {
  const data = fs.readFileSync('input.txt', 'utf8');
  const groups = data.split(/\n\n/);
  const g = groups.map(group => {
    const people = group.split(/\n/).map(x => x.trim().split("")).filter(x => x.length > 0);
    console.log(people)
    const unique = people.reduce((acc, p) => ([...new Set([...acc, ...p])]), [])
    const everyUnique = unique.filter(x => people.every(p => p.includes(x)))

    console.log(unique, everyUnique)
    return ({ group, people, unique, everyUnique });
  });
  const result1 = g.reduce((acc, { unique }) => acc + unique.length, 0);
  console.log(`result1 = ${result1}`);

  const result2 = g.reduce((acc, { everyUnique }) => acc + everyUnique.length, 0);
  console.log(`result2 = ${result2}`);
} catch (e) {
  console.error(e);
}
