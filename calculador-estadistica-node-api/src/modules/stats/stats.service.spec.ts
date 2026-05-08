import { StatsService } from "./stats.service";

describe("StatsService", () => {
  let service: StatsService;

  beforeEach(() => {
    service = new StatsService();
  });

  it("Debería calcular las estadísticas correctamente", () => {
    const data = {
      matrixQ: [
        [1, 0],
        [0, 1],
      ],
      matrixR: [
        [2, 3],
        [0, 5],
      ],
    };

    const result = service.calculateStats(data);

    expect(result.sum).toBe(12);
    expect(result.average).toBe(1.5);
    expect(result.isDiagonal).toBe(true);
  });

  it("Debería retornar false si ninguna matriz es diagonal", () => {
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
