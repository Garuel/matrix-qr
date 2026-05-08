import { StatsService } from "./stats.service";

describe("StatsService", () => {
  let service: StatsService;

  beforeEach(() => {
    service = new StatsService();
  });

  it("should calculate stats correctly for a simple matrix", () => {
    const data = {
      matrixQ: [
        [1, 0],
        [0, 1],
      ], // Diagonal
      matrixR: [
        [2, 3],
        [0, 5],
      ], // No diagonal
    };

    const result = service.calculateStats(data);

    expect(result.sum).toBe(11); // 1+0+0+1 + 2+3+0+5 = 12? No, suma manual: 1+1+2+3+5 = 12
    expect(result.average).toBe(1.5); // 12 / 8 elementos = 1.5
    expect(result.isDiagonal).toBe(true); // Porque Q es diagonal
  });

  it("should return isDiagonal as false if neither matrix is diagonal", () => {
    const data = {
      matrixQ: [
        [1, 1],
        [1, 1],
      ],
      matrixR: [
        [2, 3],
        [4, 5],
      ],
    };
    const result = service.calculateStats(data);
    expect(result.isDiagonal).toBe(false);
  });
});
