using DevicesAPI.Services;
using Microsoft.AspNetCore.Mvc;

namespace DevicesAPI.Controllers;

[Route("api/sensor")]
[ApiController]
public class SensorController: ControllerBase
{
    private readonly ISensorService _sensorService;
    private readonly ILogger<SensorController> _logger;
    
    public SensorController(ISensorService sensorService, ILogger<SensorController> logger)
    {
        _sensorService = sensorService;
        _logger = logger;
    }
    
    
    [HttpPost]
    public async Task<IActionResult> AddSensorAsync(Sensor sensorToCreate)
    {
        try
        {
            var sensor = await _sensorService.AddSensorAsync(sensorToCreate);
            return Ok(sensor);
        }
        catch (Exception ex)
        {
            _logger.LogError(ex.Message);
            return StatusCode(StatusCodes.Status500InternalServerError, ex.Message);
        }
    }
    
    [HttpGet("GetAllSensors")]
    public async Task<IActionResult> GetAllSensorsAsync()
    {
        try
        {
            IEnumerable<Sensor> sensors = await _sensorService.GetAllSensorsAsync();
            return Ok(sensors);
        }
        catch (Exception ex)
        {
            _logger.LogError(ex.Message);
            return StatusCode(StatusCodes.Status500InternalServerError, ex.Message);
        }
    }
    
    [HttpGet("GetSensorById")]
    [HttpGet("{id:int}")]
    public async Task<IActionResult> GetSensorByIdAsync(int id)
    {
        _logger.LogInformation($"Поиск сенсора с id: {id}");

        try
        {
            Sensor? sensor = await _sensorService.FindSensorByIdAsync(id);
            if (sensor == null) return NotFound();
            return Ok(sensor);
        }
        catch (Exception ex)
        {
            _logger.LogError(ex.Message);
            return StatusCode(StatusCodes.Status500InternalServerError, ex.Message);
        }
    }
}