using DevicesAPI.Models;
using Microsoft.EntityFrameworkCore;

namespace DevicesAPI.Services;


public interface ISensorService
{
    Task<Sensor> AddSensorAsync(Sensor sensorToCreate);
    // Task UpdatePersonAsync(UpdatePersonDTO personToUpdate);
    // Task DeletePersonAsync(Person person);
    Task<Sensor?> FindSensorByIdAsync(int id);
    Task<IEnumerable<Sensor>> GetAllSensorsAsync();
}
public class SensorService: ISensorService
{
    private readonly ApplicationDbContext _context;

    public SensorService(ApplicationDbContext context)
    {
        _context = context;
    }
    public async Task<Sensor> AddSensorAsync(Sensor sensorTorCreate)
    {
        _context.Sensors.Add(sensorTorCreate);
        await _context.SaveChangesAsync();
        return sensorTorCreate;
    }
    
    public async Task<Sensor?> FindSensorByIdAsync(int id)
    {
        var sensor = await _context.Sensors.Where(x => x.Id == id).AsNoTracking().FirstOrDefaultAsync();
        return sensor;
    }
    
    public async Task<IEnumerable<Sensor>> GetAllSensorsAsync()
    {
        var sensor = await _context.Sensors.AsNoTracking().ToListAsync();
        return sensor.AsEnumerable();
    }
}