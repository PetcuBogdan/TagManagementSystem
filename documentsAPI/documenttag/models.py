from django.db import models

class Document(models.Model):
    name = models.CharField(max_length=256)
    tag = models.CharField(max_length=256)

    def __str__(self):
        return self.name + " " + self.tag
        
