using System.ComponentModel;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text.Json.Serialization;

[Table("sensors")]
public class Sensor
{
    [Key]
    [Column("id")]
    public int Id { get; set; }

    [Required]
    [Column("name")]
    [MaxLength(100)]
    public string Name { get; set; }

    [Required]
    [Column("type")]
    [MaxLength(50)]
    public string Type { get; set; }

    [Required]
    [Column("location")]
    [MaxLength(100)]
    public string Location { get; set; }

    [Column("value")]
    public double Value { get; set; } = 0;

    [Column("unit")]
    [MaxLength(20)]
    public string? Unit { get; set; }

    [Required]
    [Column("status")]
    [MaxLength(20)]
    public string Status { get; set; } = "inactive";

    [Required]
    [Column("last_updated")]
    [JsonPropertyName("last_updated")]
    public DateTime LastUpdated { get; set; } = DateTime.UtcNow;

    [Required]
    [Column("created_at")]
    [JsonPropertyName("created_at")]
    public DateTime CreatedAt { get; set; } = DateTime.UtcNow;
}