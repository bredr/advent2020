import * as fs from "fs";

const ticketFn = (x: string) => {
  const row = x.slice(0, 7).split("").map(y => y === "F" ? 0 : 1).join("");
  const col = x.slice(7).split("").map(y => y === "L" ? 0 : 1).join("");
  return { row: parseInt(row, 2), col: parseInt(col, 2), sid: parseInt(row, 2) * 8 + parseInt(col, 2) }
};

try {

  console.log('BFFFBBFRRR', ticketFn('BFFFBBFRRR'));
  console.log('FFFBBBFRRR', ticketFn('FFFBBBFRRR'));
  console.log('BBFFBBFRLL', ticketFn('BBFFBBFRLL'))
  const data = fs.readFileSync('input.txt', 'utf8');
  const lines = data.split(/\r?\n/);
  const tickets = lines.map(ticketFn);
  const sorted = tickets.filter(({ sid }) => !isNaN(sid)).sort(({ sid: a }, { sid: b }) => a - b);
  console.log(sorted[sorted.length - 1])
  const rows = Math.max(...sorted.map(({ row }) => row));
  console.log(rows);
  for (let r = 1; r <= rows; r++) {
    const seats = sorted.filter(({ row }) => row === r).map(({ col }) => col);
    if (seats.length > 0) {
      let row = new Array(Math.max(...seats)).fill("E");
      seats.forEach(i => { row[i] = "X"; });
      if (row.join("").indexOf("XEX") > -1) {
        console.log(r, seats)
      }
    }
  }

} catch (err) {
  console.error(err)
}

