import json
from django.http import JsonResponse
from .serializers import DocumentSerializer
from rest_framework.decorators import api_view
from rest_framework.response import Response
from rest_framework import status
from .models import Document
from rest_framework.permissions import AllowAny
from rest_framework.views import APIView
#from rest_framework.filters import SearchFilter, OrderingFilter
#from documenttag import serializers

@api_view(['GET', 'POST'])
def document_list(request):
    if request.method == 'GET':
        documents = Document.objects.all()
        serializer = DocumentSerializer(documents, many=True)
        return JsonResponse(serializer.data, safe=False)
    if request.method == 'POST':
        serializer = DocumentSerializer(data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_201_CREATED)

@api_view(['GET', 'POST', 'DELETE'])
def document_detail(request, id_document):

    try:
        document  = Document.objects.get(pk=id_document)
    except Document.DoesNotExist():
        return Response(status=status.HTTP_404_NOT_FOUND)

    if request.method == 'GET':
        serializer = DocumentSerializer(document)
        return JsonResponse(serializer.data)
    if request.method == 'POST':
        serializer = DocumentSerializer(document, data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data)
        #return Response(serializer.errors(), status=status.HTTP_400_BAD_REQUEST)
    if request.method == 'DELETE':
        document.delete()
        return Response(status=status.HTTP_204_NO_CONTENT)

class Search(APIView):
    permission_classes = [AllowAny]

    def search_document(request):
        if request.method == 'POST':
            data = json.loads(request.body.decode('utf-8'))
            searched = data.get('tag')
            documents = Document.objects.filter(tag__contains=searched)
            serializer = DocumentSerializer(documents, many=True)
            return JsonResponse(serializer.data, safe=False)
# Create your views here.
