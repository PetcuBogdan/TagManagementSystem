from dataclasses import field
from xml.dom.minidom import Document
from rest_framework import serializers
from .models import Document

class DocumentSerializer(serializers.ModelSerializer):
    class Meta:
        model = Document
        fields = ['name', 'tag']
        