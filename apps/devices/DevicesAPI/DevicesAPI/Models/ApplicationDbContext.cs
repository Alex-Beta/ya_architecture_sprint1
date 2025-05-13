

using DevicesAPI.Models;
using Microsoft.EntityFrameworkCore;

namespace DevicesAPI.Models;

public class ApplicationDbContext : DbContext
{
    public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options) : base(options)
    {

    }

    public DbSet<Sensor> Sensors { get; set; }
}